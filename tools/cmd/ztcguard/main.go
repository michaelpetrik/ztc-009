package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type event struct {
	Session string  `json:"session_id"`
	Turn    string  `json:"turn_id"`
	CWD     string  `json:"cwd"`
	Prompt  string  `json:"prompt"`
	Message *string `json:"last_assistant_message"`
}

type state struct {
	Enabled   bool            `json:"enabled"`
	Turn      int             `json:"turn"`
	HostTurn  string          `json:"host_turn,omitempty"`
	Toggled   int             `json:"toggled"`
	Mode      string          `json:"mode"`
	Attempts  int             `json:"attempts"`
	Repair    string          `json:"repair,omitempty"`
	Trusted   map[string]bool `json:"trusted,omitempty"`
	AllowCode bool            `json:"allow_code,omitempty"`
}

func stateFile(session string) (string, error) {
	if session == "" {
		return "", fmt.Errorf("chybí session_id")
	}
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "ztc-009", "sessions", fmt.Sprintf("%x.json", sha256.Sum256([]byte(session)))), nil
}

func transaction(path string, apply func(*state, bool) (any, error)) (any, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}
	lock := path + ".lock"
	deadline := time.Now().Add(time.Second)
	for {
		err := os.Mkdir(lock, 0700)
		if err == nil {
			break
		}
		if !os.IsExist(err) || time.Now().After(deadline) {
			return nil, fmt.Errorf("stav relace je zamčený: %w", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	defer os.Remove(lock)
	var s state
	data, err := os.ReadFile(path)
	exists := err == nil
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if exists {
		if err := json.Unmarshal(data, &s); err != nil {
			return nil, fmt.Errorf("poškozený stav relace: %w", err)
		}
	}
	out, err := apply(&s, exists)
	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(s)
	if err != nil {
		return nil, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".state-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return nil, err
	}
	if err := tmp.Close(); err != nil {
		return nil, err
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return nil, err
	}
	return out, nil
}

func initial(cwd string) (bool, error) {
	if cwd == "" {
		return false, nil
	}
	b, err := os.ReadFile(filepath.Join(cwd, ".ztc", "profile.md"))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return regexp.MustCompile(`(?m)^enabled: true\s*$`).Match(b) && regexp.MustCompile(`(?m)^scope_confirmed: true\s*$`).Match(b), nil
}

func handle(action string, e event, s *state, exists bool, executable string) (any, error) {
	switch action {
	case "prompt":
		if !exists {
			enabled, err := initial(e.CWD)
			if err != nil {
				return nil, err
			}
			s.Enabled = enabled
		}
		continuation := s.Repair != "" && e.Prompt == s.Repair
		if continuation && e.Turn != "" {
			s.HostTurn = e.Turn
		}
		duplicate := e.Turn != "" && s.HostTurn == e.Turn
		if !continuation && !duplicate {
			s.Turn++
			s.HostTurn, s.Mode, s.Repair, s.Attempts = e.Turn, "enforce", "", 0
			s.Trusted = map[string]bool{}
			_, spans, _ := prose(e.Prompt, nil, false)
			for _, span := range spans {
				s.Trusted[span] = true
			}
			s.AllowCode = false
			command := strings.TrimSpace(strings.SplitN(e.Prompt, "\n", 2)[0])
			if strings.HasPrefix(command, "ztc ") {
				command = "/" + command
			}
			switch command {
			case "/ztc skip":
				s.Mode = "outside"
			case "/ztc once":
				s.Mode = "force"
			case "/ztc code":
				s.AllowCode = true
			}
			switch strings.ToLower(strings.TrimSpace(e.Prompt)) {
			case "ahoj", "ahoj!", "díky", "děkuji", "děkuju":
				s.Mode = "trivial"
			}
			switch strings.TrimSpace(e.Prompt) {
			case "/ztc on":
				s.Enabled = true
				s.Toggled = s.Turn
			case "/ztc off":
				s.Enabled = false
				s.Toggled = s.Turn
			case "/ztc", "/ztc-009:ztc", "$ztc", "/skill:ztc", "ztc toggle":
				s.Enabled = !s.Enabled
				s.Toggled = s.Turn
			}
		}
		context := fmt.Sprintf("ZTČ-GUARD: session_enabled=%t; turn=%d; mode=%s. Stav řídí uživatelský příkaz, nikoli model. ", s.Enabled, s.Turn, s.Mode)
		context += "Nepoužívej shell k přepínání ani simulaci hooků. Samostatný /ztc přepne relaci. Pokud klient rozbalil skill a přepnutí nenastalo, požádej uživatele o běžnou zprávu ztc toggle. "
		context += "Uživatel může zadat /ztc skip na prvním řádku pro jednu odpověď mimo režim, /ztc once pro jednorázový převod a /ztc code pro povolení nového kódu. Sám tyto výjimky neuděluj. "
		context += "Kontrola vyžaduje nejvýše 20 slov ve větě včetně závorek a nadpisů, šest vět v odstavci, bez středníku a obecných zkratek. Vysvětli CI/VM/FE/BE/UI/DB/PR. "
		context += "Citace a kód jsou chráněné jen při shodě s podklady v aktuálním uživatelském zadání. Po chybě oprav celou odpověď bez změny významu."

		return map[string]any{"hookSpecificOutput": map[string]string{"hookEventName": "UserPromptSubmit", "additionalContext": context}}, nil
	case "toggle":
		if !exists || s.Toggled != s.Turn || s.Turn == 0 {
			return nil, errors.New("přepnutí vyžaduje uživatelský příkaz v UserPromptSubmit")
		}
		return map[string]any{"session_enabled": s.Enabled, "turn": s.Turn, "mode": s.Mode}, nil
	case "on", "off", "outside", "trivial", "enforce":
		return nil, errors.New("změna stavu z modelového nástroje není povolená; použijte uživatelský příkaz")

	case "stop":
		if !exists {
			return nil, fmt.Errorf("chybí stav z UserPromptSubmit; kontrola neproběhla")
		}
		if e.Turn != "" && s.HostTurn != "" && e.Turn != s.HostTurn {
			return nil, errors.New("Stop patří jinému tahu")
		}
		if (!s.Enabled && s.Mode != "force") || s.Mode == "outside" {
			return map[string]any{}, nil
		}
		if e.Message == nil || strings.TrimSpace(*e.Message) == "" {
			return nil, fmt.Errorf("hostitel neposkytl text odpovědi")
		}
		findings := checkProtected(*e.Message, s.Mode == "trivial", s.Trusted, s.AllowCode)
		if len(findings) == 0 {
			return map[string]any{}, nil
		}
		if s.Attempts >= 3 {
			msg := "ZTČ-009: kontrola selhala i po třech opravách. Odpověď není ověřená. " + strings.Join(findings, "; ")
			return map[string]any{"continue": false, "stopReason": msg, "systemMessage": msg}, nil
		}
		s.Attempts++
		s.Repair = "[ZTČ-GUARD] Oprav celou poslední odpověď. Zachovej význam, kód a citace. Neobcházej kontrolu výjimkou. Chyby: " + strings.Join(findings, "; ")
		return map[string]string{"decision": "block", "reason": s.Repair}, nil
	}
	return nil, fmt.Errorf("neznámá akce: %s", action)
}

func run(args []string, input io.Reader) (any, error) {
	flags := flag.NewFlagSet("ztcguard", flag.ContinueOnError)
	session := flags.String("session", "", "session ID supplied by the hook")
	turn := flags.Int("turn", 0, "turn supplied by the hook")
	action := flags.String("action", "", "prompt, stop, toggle, on, off, enforce, outside, trivial")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	if flags.NArg() != 0 {
		return nil, errors.New("unexpected positional arguments")
	}
	e := event{Session: *session}
	if *action == "prompt" || *action == "stop" {
		data, err := io.ReadAll(io.LimitReader(input, 2*1024*1024+1))
		if err != nil {
			return nil, err
		}
		if len(data) > 2*1024*1024 {
			return nil, errors.New("hook input too large")
		}
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
	}
	path, err := stateFile(e.Session)
	if err != nil {
		return nil, err
	}
	executable, err := os.Executable()
	if err != nil {
		return nil, err
	}
	return transaction(path, func(s *state, exists bool) (any, error) {
		if *action != "prompt" && *action != "stop" && (*turn == 0 || *turn != s.Turn) {
			return nil, errors.New("stale or missing turn")
		}
		out, err := handle(*action, e, s, exists, executable)
		return out, err
	})
}

func main() {
	if len(os.Args) == 3 && os.Args[1] == "--action" && os.Args[2] == "validate-output" {
		data, err := io.ReadAll(io.LimitReader(os.Stdin, 65537))
		if err != nil {
			os.Exit(1)
		}
		result, err := validateOutput(data)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("ZTC-GUARD/1 %s\n", result)
		return
	}

	out, err := run(os.Args[1:], os.Stdin)
	if err != nil {
		msg := "ZTČ-009: kontrola není funkční: " + err.Error()
		out = map[string]any{"continue": false, "stopReason": msg, "systemMessage": msg}
	}
	if err := json.NewEncoder(os.Stdout).Encode(out); err != nil {
		os.Exit(1)
	}
}

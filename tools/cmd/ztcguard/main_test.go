package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func invoke(t *testing.T, action string, e event, s *state, exists bool) map[string]any {
	t.Helper()
	out, err := handle(action, e, s, exists, "/path with spaces/guard")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(out)
	var result map[string]any
	if err := json.Unmarshal(b, &result); err != nil {
		t.Fatal(err)
	}
	return result
}

func TestToggleAndIsolation(t *testing.T) {
	var a, b state
	e := event{Session: "a", Turn: "turn1", Prompt: "/ztc"}
	invoke(t, "prompt", e, &a, false)
	if !a.Enabled {
		t.Fatal("first toggle must enable")
	}
	invoke(t, "toggle", e, &a, true)
	invoke(t, "prompt", e, &a, true)
	if !a.Enabled || a.Turn != 1 {
		t.Fatal("same turn toggled twice")
	}
	invoke(t, "prompt", event{Session: "b", Prompt: "Dotaz"}, &b, false)
	if b.Enabled {
		t.Fatal("session leaked")
	}
	e.Turn = "turn2"
	invoke(t, "prompt", e, &a, true)
	if a.Enabled {
		t.Fatal("second command must disable")
	}
	invoke(t, "prompt", event{Prompt: "/ztc on", Turn: "three"}, &a, true)
	if !a.Enabled {
		t.Fatal("explicit on failed")
	}
	invoke(t, "prompt", event{Prompt: "/ztc off", Turn: "four"}, &a, true)
	if a.Enabled {
		t.Fatal("explicit off failed")
	}
	invoke(t, "prompt", event{Prompt: "Vysvětli příkaz /ztc"}, &b, true)
	if b.Enabled {
		t.Fatal("quoted command toggled")
	}
}

func TestRepairCycle(t *testing.T) {
	s := state{Enabled: true, Turn: 1, Mode: "enforce"}
	bad := "Systém běží; chyba trvá."
	for i := 0; i < 3; i++ {
		out := invoke(t, "stop", event{Message: &bad}, &s, true)
		if out["decision"] != "block" {
			t.Fatalf("missing block: %v", out)
		}
		invoke(t, "prompt", event{Prompt: s.Repair, Turn: "continuation"}, &s, true)
		if s.Attempts != i+1 {
			t.Fatal("continuation reset retry budget")
		}
	}
	out := invoke(t, "stop", event{Message: &bad}, &s, true)
	if out["continue"] != false || out["systemMessage"] == nil {
		t.Fatalf("must report failure and stop: %v", out)
	}
	good := "Systém běží. Chyba trvá."
	if out := invoke(t, "stop", event{Message: &good}, &s, true); len(out) != 0 {
		t.Fatalf("valid correction blocked: %v", out)
	}
	invoke(t, "prompt", event{Prompt: "Nový dotaz", Turn: "new"}, &s, true)
	if s.Attempts != 0 || s.Mode != "enforce" {
		t.Fatal("new prompt did not reset turn")
	}
}

func TestModesAndMissingInput(t *testing.T) {
	bad := "CI; další věta."
	s := state{Turn: 1}
	if out := invoke(t, "stop", event{Message: &bad}, &s, true); len(out) != 0 {
		t.Fatal("disabled mode enforced")
	}
	invoke(t, "prompt", event{Prompt: "/ztc once\nPřeveď text."}, &s, true)
	if out := invoke(t, "stop", event{Message: &bad}, &s, true); out["decision"] != "block" {
		t.Fatal("one-off conversion bypassed")
	}
	if _, err := handle("outside", event{}, &s, true, ""); err == nil {
		t.Fatal("repair bypass accepted")
	}
	s.Attempts = 0
	invoke(t, "prompt", event{Prompt: "/ztc skip\nJiné téma."}, &s, true)
	if out := invoke(t, "stop", event{Message: &bad}, &s, true); len(out) != 0 {
		t.Fatal("outside scope blocked")
	}
	s.Enabled = true
	invoke(t, "prompt", event{Prompt: "ahoj"}, &s, true)
	one := "Ano, CI běží."
	if out := invoke(t, "stop", event{Message: &one}, &s, true); len(out) != 0 {
		t.Fatal("trivial answer blocked")
	}
	if _, err := handle("stop", event{}, &s, true, ""); err == nil {
		t.Fatal("missing response accepted")
	}
	if _, err := handle("stop", event{}, &s, false, ""); err == nil {
		t.Fatal("missing state accepted")
	}
}

func TestModelCannotGrantExceptions(t *testing.T) {
	for _, action := range []string{"toggle", "on", "off", "outside", "trivial", "enforce"} {
		s := state{Enabled: true, Turn: 1, Mode: "enforce"}
		if _, err := handle(action, event{}, &s, true, ""); err == nil {
			t.Errorf("model action %s accepted", action)
		}
		if !s.Enabled || s.Mode != "enforce" {
			t.Errorf("action %s changed state", action)
		}
	}
	s := state{Enabled: true, Turn: 2, HostTurn: "new", Mode: "enforce"}
	text := "Systém běží."
	if _, err := handle("stop", event{Turn: "old", Message: &text}, &s, true, ""); err == nil {
		t.Fatal("stale Stop accepted")
	}
}

func TestStateIsAtomicAndPrivate(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sessions", "a.json")
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := transaction(path, func(s *state, _ bool) (any, error) { s.Turn++; return nil, nil })
			if err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var s state
	if err := json.Unmarshal(b, &s); err != nil || s.Turn != 12 {
		t.Fatalf("lost update: %s %v", b, err)
	}
	info, _ := os.Stat(path)
	if info.Mode().Perm() != 0600 {
		t.Fatal("state not private")
	}
	if err := os.WriteFile(path, []byte("broken"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := transaction(path, func(s *state, _ bool) (any, error) { return nil, nil }); err == nil {
		t.Fatal("corrupt state silently reset")
	}
}

func TestProfileAndSafeIDs(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, ".ztc"), 0700); err != nil {
		t.Fatal(err)
	}
	profile := filepath.Join(dir, ".ztc", "profile.md")
	content := []byte("enabled: true\nscope_confirmed: true\n")
	if err := os.WriteFile(profile, content, 0600); err != nil {
		t.Fatal(err)
	}
	var s state
	invoke(t, "prompt", event{CWD: dir, Prompt: "/ztc"}, &s, false)
	if s.Enabled {
		t.Fatal("toggle ignored active profile")
	}
	b, _ := os.ReadFile(profile)
	if string(b) != string(content) {
		t.Fatal("profile changed")
	}
	p, err := stateFile("../../private;$(touch injected)")
	if err != nil || strings.Contains(filepath.Base(p), "private") {
		t.Fatal("unsafe session filename")
	}
	if _, err := stateFile(""); err == nil {
		t.Fatal("missing session accepted")
	}
}

func TestWireProtocol(t *testing.T) {
	session := "ztcguard-test-" + t.TempDir()
	path, err := stateFile(session)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Remove(path); os.Remove(path + ".lock") })
	prompt, _ := json.Marshal(event{Session: session, Prompt: "/ztc", Turn: "test-turn"})
	if _, err := run([]string{"--action", "prompt"}, strings.NewReader(string(prompt))); err != nil {
		t.Fatal(err)
	}
	if _, err := run([]string{"--action", "toggle", "--session", session, "--turn", "1"}, strings.NewReader("")); err != nil {
		t.Fatal(err)
	}
	if _, err := run([]string{"--action", "outside", "--session", session, "--turn", "999"}, strings.NewReader("")); err == nil {
		t.Fatal("stale command accepted")
	}
	bad := "Systém běží; chyba trvá."
	stop, _ := json.Marshal(event{Session: session, Message: &bad})
	out, err := run([]string{"--action", "stop"}, strings.NewReader(string(stop)))
	if err != nil {
		t.Fatal(err)
	}
	decision, ok := out.(map[string]string)
	if !ok || decision["decision"] != "block" {
		t.Fatalf("bad output passed wire protocol: %v", out)
	}
	good := "Systém běží. Chyba trvá."
	stop, _ = json.Marshal(event{Session: session, Message: &good})
	out, err = run([]string{"--action", "stop"}, strings.NewReader(string(stop)))
	if err != nil || len(out.(map[string]any)) != 0 {
		t.Fatalf("correction failed: %v %v", out, err)
	}
	if _, err := run([]string{"--action", "prompt"}, strings.NewReader("not json")); err == nil {
		t.Fatal("malformed event accepted")
	}
}

func TestLauncherEndToEnd(t *testing.T) {
	dir := t.TempDir()
	binary := filepath.Join(dir, "guard with spaces")
	if out, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v %s", err, out)
	}
	session := "ztcguard-launcher-" + dir
	path, err := stateFile(session)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Remove(path); os.Remove(path + ".lock") })
	launch := func(action string, e event) map[string]any {
		t.Helper()
		payload, _ := json.Marshal(e)
		cmd := exec.Command("sh", "../../../hooks/ztc-guard.sh", action)
		for _, item := range os.Environ() {
			if !strings.HasPrefix(item, "ZTC_GUARD_BIN=") {
				cmd.Env = append(cmd.Env, item)
			}
		}
		cmd.Env = append(cmd.Env, "ZTC_GUARD_BIN="+binary)
		cmd.Stdin = strings.NewReader(string(payload))
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("launcher: %v %s", err, out)
		}
		var result map[string]any
		if err := json.Unmarshal(out, &result); err != nil {
			t.Fatalf("JSON: %v %s", err, out)
		}
		return result
	}
	launch("prompt", event{Session: session, Prompt: "/ztc", Turn: "one"})
	bad := "Systém běží; chyba trvá."
	if result := launch("stop", event{Session: session, Message: &bad}); result["decision"] != "block" {
		t.Fatalf("bad answer passed: %v", result)
	}
	good := "Systém běží. Chyba trvá."
	if result := launch("stop", event{Session: session, Message: &good}); len(result) != 0 {
		t.Fatalf("good answer blocked: %v", result)
	}
	launch("prompt", event{Session: session, Prompt: "/ztc", Turn: "two"})
	if result := launch("stop", event{Session: session, Message: &bad}); len(result) != 0 {
		t.Fatalf("disabled session blocked: %v", result)
	}
	binary = filepath.Join(dir, "crashed guard")
	if err := os.WriteFile(binary, []byte("#!/bin/sh\nexit 1\n"), 0700); err != nil {
		t.Fatal(err)
	}
	if result := launch("stop", event{Session: session, Message: &good}); result["continue"] != false {
		t.Fatalf("crash passed: %v", result)
	}
	for _, body := range []string{"#!/bin/sh\nexit 0\n", "#!/bin/sh\nprintf 'invalid JSON'\n", "#!/bin/sh\nprintf '{}'\n"} {
		if err := os.WriteFile(binary, []byte(body), 0700); err != nil {
			t.Fatal(err)
		}
		if result := launch("stop", event{Session: session, Message: &good}); result["continue"] != false {
			t.Fatalf("invalid protocol passed: %v", result)
		}
	}
}

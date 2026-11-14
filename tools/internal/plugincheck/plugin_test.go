package plugincheck

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func root(t *testing.T) string {
	t.Helper()
	p, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func readJSON(t *testing.T, name string, target any) {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root(t), name))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, target); err != nil {
		t.Fatal(err)
	}
}

type hook struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Timeout int    `json:"timeout"`
}

func configuredHook(t *testing.T) hook {
	t.Helper()
	var config struct {
		Hooks map[string][]struct {
			Hooks []hook `json:"hooks"`
		} `json:"hooks"`
	}
	readJSON(t, "hooks/hooks.json", &config)
	events := config.Hooks["UserPromptSubmit"]
	if len(config.Hooks) != 1 || len(events) != 1 || len(events[0].Hooks) != 1 {
		t.Fatal("expected one UserPromptSubmit hook")
	}
	h := events[0].Hooks[0]
	if h.Type != "command" || h.Command == "" || h.Timeout < 1 || h.Timeout > 5 {
		t.Fatalf("invalid hook: %+v", h)
	}
	return h
}

func TestManifests(t *testing.T) {
	for _, name := range []string{".claude-plugin/plugin.json", ".codex-plugin/plugin.json"} {
		var manifest struct {
			Name        string `json:"name"`
			Version     string `json:"version"`
			Description string `json:"description"`
			Skills      string `json:"skills"`
			Hooks       string `json:"hooks"`
		}
		readJSON(t, name, &manifest)
		if manifest.Name != "ztc-009" || manifest.Version == "" || manifest.Description == "" {
			t.Fatalf("incomplete manifest %s", name)
		}
		if strings.HasPrefix(name, ".codex-plugin") && (manifest.Skills != "./skills/" || manifest.Hooks != "./hooks/hooks.json") {
			t.Fatal("Codex manifest must reference shared skills and hooks")
		}
	}
	configuredHook(t)
}

func TestHookIsReadOnlyAndDoesNotReflectPrompt(t *testing.T) {
	plugin := filepath.Join(t.TempDir(), "plugin with spaces")
	for _, path := range []string{"hooks", "skills/ztc-response"} {
		if err := os.MkdirAll(filepath.Join(plugin, path), 0755); err != nil {
			t.Fatal(err)
		}
	}
	script, err := os.ReadFile(filepath.Join(root(t), "hooks/ztc-context.sh"))
	if err != nil {
		t.Fatal(err)
	}
	files := map[string][]byte{
		"hooks/ztc-context.sh":         script,
		"skills/ztc-response/SKILL.md": []byte("test skill"),
	}
	for name, data := range files {
		if err := os.WriteFile(filepath.Join(plugin, name), data, 0644); err != nil {
			t.Fatal(err)
		}
	}
	before := snapshot(t, plugin)
	command := configuredHook(t).Command
	run := func(input string) []byte {
		t.Helper()
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = plugin
		cmd.Env = append(cleanEnv(), "CLAUDE_PLUGIN_ROOT="+plugin)
		cmd.Stdin = strings.NewReader(input)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("hook failed: %v: %s", err, out)
		}
		return out
	}
	first := run(`{"prompt":"PRIVATE_PROMPT_MARKER $(touch injected)"}`)
	second := run("not JSON")
	if len(first) == 0 || string(first) != string(second) || strings.Contains(string(first), "PRIVATE_PROMPT_MARKER") {
		t.Fatal("hook must emit static context independent of stdin")
	}
	if !strings.Contains(string(first), `"`+filepath.Join(plugin, "skills/ztc-response/SKILL.md")+`"`) {
		t.Fatal("hook must provide the quoted absolute skill path")
	}
	if after := snapshot(t, plugin); !reflect.DeepEqual(before, after) {
		t.Fatal("hook changed plugin state")
	}
}

func TestMissingRootIsFailSoft(t *testing.T) {
	for _, value := range []string{"", filepath.Join(t.TempDir(), "missing")} {
		cmd := exec.Command("sh", filepath.Join(root(t), "hooks/ztc-context.sh"))
		cmd.Env = append(cleanEnv(), "CLAUDE_PLUGIN_ROOT="+value)
		out, err := cmd.CombinedOutput()
		if err != nil || len(out) != 0 {
			t.Fatalf("missing root must succeed silently: %v %s", err, out)
		}
	}
}

func cleanEnv() []string {
	var env []string
	for _, entry := range os.Environ() {
		if !strings.HasPrefix(entry, "CLAUDE_PLUGIN_ROOT=") {
			env = append(env, entry)
		}
	}
	return env
}

func snapshot(t *testing.T, dir string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			files[path] = "directory"
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		files[path] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ztc-009/ztc-009/tools/internal/dictionary"
)

func TestDiscoveredDictionaryValidation(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "slovnik")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}
	header := "lemma,pos,vid,status,meaning,replace_with,en_ste,ok,bad\n"
	for _, name := range []string{"ztc-009.csv", "healthcare.csv"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(header+"soubor,N,,approved,jednotka dat,,,Zkontrolujte soubor.,\n"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	var software *dictionary.Checker
	checks, err := checkers(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, checker := range checks {
		if c, ok := checker.(*dictionary.Checker); ok && c.Path == filepath.Join(dir, "healthcare.csv") {
			software = c
		}
	}
	if software == nil {
		t.Fatal("software dictionary is not registered")
	}
	if issues, err := software.Check(); err != nil || len(issues) != 0 {
		t.Fatalf("valid dictionary: %v %v", issues, err)
	}
	if err := os.WriteFile(software.Path, []byte(header+"soubor,N,,approved,,,,Zkontrolujte soubor.,\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if issues, err := software.Check(); err != nil || len(issues) != 1 || issues[0].Rule != "status" {
		t.Fatalf("missing meaning: %v %v", issues, err)
	}
	if err := os.Remove(software.Path); err != nil {
		t.Fatal(err)
	}
	if _, err := software.Check(); err == nil {
		t.Fatal("missing software dictionary must fail")
	}
}

func TestDiscoveryErrors(t *testing.T) {
	root := t.TempDir()
	if _, err := checkers(root); err == nil {
		t.Fatal("missing dictionary directory must fail")
	}
	if err := os.Mkdir(filepath.Join(root, "slovnik"), 0755); err != nil {
		t.Fatal(err)
	}
	if _, err := checkers(root); err == nil {
		t.Fatal("empty directory must fail")
	}
	if err := os.WriteFile(filepath.Join(root, "slovnik", "healthcare.csv"), []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := checkers(root); err == nil {
		t.Fatal("missing base dictionary must fail")
	}
}

func TestStandaloneDictionary(t *testing.T) {
	dir := t.TempDir()
	binary := filepath.Join(dir, "ztcvalidate")
	if output, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %s: %v", output, err)
	}
	path := filepath.Join(dir, "working.csv")
	cases := []struct {
		name, content string
		code          int
	}{
		{"valid", "lemma,pos,vid,status,meaning,replace_with,en_ste,ok,bad\nsoubor,N,,approved,jednotka dat,,,Zkontrolujte soubor.,\n", 0},
		{"invalid", "lemma,pos,vid,status,meaning,replace_with,en_ste,ok,bad\nsoubor,N,,approved,,,,Zkontrolujte soubor.,\n", 1},
		{"missing", "", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "missing" {
				if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
					t.Fatal(err)
				}
			} else if err := os.WriteFile(path, []byte(tc.content), 0644); err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command(binary, "-root", filepath.Join(dir, "absent-repo"), "-dictionary", path)
			cmd.Dir = dir
			output, err := cmd.CombinedOutput()
			code := 0
			if err != nil {
				exit, ok := err.(*exec.ExitError)
				if !ok {
					t.Fatal(err)
				}
				code = exit.ExitCode()
			}
			if code != tc.code {
				t.Fatalf("exit %d, want %d: %s", code, tc.code, output)
			}
		})
	}
}

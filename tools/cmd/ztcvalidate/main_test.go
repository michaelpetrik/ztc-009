package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ztc-009/ztc-009/tools/internal/dictionary"
)

func TestSoftwareDictionaryValidation(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "slovnik")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}
	header := "lemma,pos,vid,status,meaning,replace_with,en_ste,ok,bad\n"
	for _, name := range []string{"ztc-009.csv", "software-development.csv"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(header+"soubor,N,,approved,jednotka dat,,,Zkontrolujte soubor.,\n"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	var software *dictionary.Checker
	for _, checker := range checkers(root) {
		if c, ok := checker.(*dictionary.Checker); ok && c.Path == filepath.Join(dir, "software-development.csv") {
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

package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, root, name, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	for _, name := range required {
		write(t, root, name, name)
	}
	return root
}

func TestPackageResourcesAndExclusions(t *testing.T) {
	root := fixture(t)
	write(t, root, ".gitignore", "skills/credentials.json\n")
	wanted := []string{".claude-plugin/marketplace.json", ".agents/plugins/marketplace.json", "skills/ztc-install/references/profile-template.md", "slovnik/software.csv", "docs/HARNESSES.md", "extensions/ztc.js", "tools/cmd/ztcvalidate/main.go", "tools/go.mod", "tools/LICENSE"}
	for _, name := range wanted {
		write(t, root, name, "resource: "+name)
	}
	excluded := []string{"skills/credentials.json", "hooks/config.local.json", "extensions/private-notes.md", ".git/config", ".ztc/profile.json", "docs/audit/private.md", "skills/.env", "skills/key.pem", "skills/private.key", "skills/tests/test.md", "skills/.private/secret.json", "extensions/pi/index.test.ts", "tools/cmd/ztcvalidate/main_test.go", "tools/token.json", "tools/ztcvalidate"}
	for _, name := range excluded {
		write(t, root, name, "SENSITIVE")
	}
	output := filepath.Join(root, "skills", "previous.zip")
	write(t, root, "skills/previous.zip", "OLD ARCHIVE")
	if err := pack(root, output); err != nil {
		t.Fatal(err)
	}
	r, err := zip.OpenReader(output)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	seen := map[string]bool{}
	for _, entry := range r.File {
		if strings.HasPrefix(entry.Name, "/") || strings.Contains(entry.Name, "..") || strings.Contains(entry.Name, "\\") {
			t.Fatalf("unsafe path %q", entry.Name)
		}
		seen[entry.Name] = true
		f, err := entry.Open()
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(body), "SENSITIVE") {
			t.Fatalf("leaked %s", entry.Name)
		}
	}
	for _, name := range append(required, wanted...) {
		if !seen[name] {
			t.Errorf("missing %s", name)
		}
	}
	for _, name := range append(excluded, "skills/previous.zip") {
		if seen[name] {
			t.Errorf("included %s", name)
		}
	}
}

func TestFailurePreservesOutput(t *testing.T) {
	for _, mode := range []string{"missing-required", "symlink-file", "symlink-directory", "output-directory"} {
		t.Run(mode, func(t *testing.T) {
			root := fixture(t)
			output := filepath.Join(t.TempDir(), "plugin.zip")
			if err := os.WriteFile(output, []byte("original"), 0644); err != nil {
				t.Fatal(err)
			}
			switch mode {
			case "missing-required":
				if err := os.Remove(filepath.Join(root, required[0])); err != nil {
					t.Fatal(err)
				}
			case "symlink-file", "symlink-directory":
				target := output
				if mode == "symlink-directory" {
					target = t.TempDir()
				}
				link := filepath.Join(root, "hooks", "ztc-context.sh")
				if mode == "symlink-directory" {
					link = filepath.Join(root, "hooks")
				} else if err := os.MkdirAll(filepath.Dir(link), 0755); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, link); err != nil {
					t.Fatal(err)
				}
			case "output-directory":
				output = filepath.Join(output, "plugin.zip")
			}
			if err := pack(root, output); err == nil {
				t.Fatal("expected error")
			}
			if mode != "output-directory" {
				body, err := os.ReadFile(output)
				if err != nil || string(body) != "original" {
					t.Fatalf("changed original: %q %v", body, err)
				}
			}
			leftovers, err := filepath.Glob(filepath.Join(filepath.Dir(output), ".ztcpack-*.zip"))
			if err != nil || len(leftovers) != 0 {
				t.Fatalf("temporary files remain: %v %v", leftovers, err)
			}
		})
	}
}

func TestOptionalFilesMayBeAbsent(t *testing.T) {
	root := fixture(t)
	if err := pack(root, filepath.Join(t.TempDir(), "new", "directory", "plugin.zip")); err != nil {
		t.Fatal(err)
	}
}

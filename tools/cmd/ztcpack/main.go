package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var directories = []string{".claude-plugin", ".codex-plugin", ".agents/plugins", "skills", "hooks", "llm", "slovnik", "spec", "extensions", "tools"}
var files = []string{"README.md", "CONTRIBUTING.md", "CHANGELOG.md", "LICENSE", "NOTICE", "package.json", "docs/INSTALL.md", "docs/CHAT.md", "docs/HARNESSES.md", "docs/ztc-flow.svg", "docs/ztc-flow.mmd"}
var required = []string{".claude-plugin/plugin.json", ".codex-plugin/plugin.json", "skills/ztc-install/SKILL.md", "skills/ztc-response/SKILL.md", "llm/ztc-009.llm.md", "slovnik/ztc-009.csv"}

func distributable(name string) bool {
	base := filepath.Base(name)
	if strings.HasPrefix(base, ".") || strings.Contains(base, ".test.") || strings.Contains(base, ".spec.") || strings.HasSuffix(base, "_test.go") {
		return false
	}
	if base == "LICENSE" || base == "NOTICE" {
		return true
	}
	switch filepath.Ext(base) {
	case ".md", ".csv", ".json", ".sh", ".ts", ".js", ".mjs", ".svg", ".mmd", ".yaml", ".yml":
		return true
	}
	return false
}

func collect(root, output string) ([]string, error) {
	var result []string
	for _, name := range required {
		info, err := os.Lstat(filepath.Join(root, filepath.FromSlash(name)))
		if err != nil {
			return nil, fmt.Errorf("required file %s: %w", name, err)
		}
		if !info.Mode().IsRegular() {
			return nil, fmt.Errorf("required file is not regular: %s", name)
		}
	}
	add := func(path string, info fs.FileInfo) error {
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink is not allowed: %s", path)
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular file is not allowed: %s", path)
		}
		if path == output {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		include := distributable(path)
		if strings.HasPrefix(filepath.ToSlash(rel), "tools/") {
			include = (strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")) || rel == filepath.Join("tools", "go.mod") || rel == filepath.Join("tools", "LICENSE")
		}
		if include {
			result = append(result, rel)
		}
		return nil
	}
	for _, directory := range directories {
		start := filepath.Join(root, filepath.FromSlash(directory))
		if _, err := os.Lstat(start); os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil, err
		}
		err := filepath.WalkDir(start, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.Type()&os.ModeSymlink != 0 {
				return fmt.Errorf("symlink is not allowed: %s", path)
			}
			if entry.IsDir() {
				if path != start && (strings.HasPrefix(entry.Name(), ".") || entry.Name() == "tests" || entry.Name() == "test" || entry.Name() == "__tests__" || entry.Name() == "audit") {
					return filepath.SkipDir
				}
				return nil
			}
			info, err := entry.Info()
			if err != nil {
				return err
			}
			return add(path, info)
		})
		if err != nil {
			return nil, err
		}
	}
	for _, name := range files {
		path := filepath.Join(root, filepath.FromSlash(name))
		info, err := os.Lstat(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if err := add(path, info); err != nil {
			return nil, err
		}
	}
	sort.Strings(result)
	return result, nil
}

func pack(root, output string) error {
	root, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	output, err = filepath.Abs(output)
	if err != nil {
		return err
	}
	paths, err := collect(root, output)
	if err != nil {
		return err
	}
	if filepath.Ext(output) != ".zip" {
		return fmt.Errorf("output must have .zip extension")
	}
	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(output), ".ztcpack-*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()
	archive := zip.NewWriter(tmp)
	for _, path := range paths {
		if err := appendFile(archive, root, path); err != nil {
			archive.Close()
			return err
		}
	}
	if err := archive.Close(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), output)
}

func appendFile(archive *zip.Writer, root, path string) error {
	file, err := os.Open(filepath.Join(root, path))
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.ToSlash(path)
	header.Method = zip.Deflate
	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	return err
}

func main() {
	root := flag.String("root", ".", "repository root")
	output := flag.String("out", "ztc-009.plugin.zip", "output ZIP path")
	flag.Parse()
	if err := pack(*root, *output); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(*output)
}

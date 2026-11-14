package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

var files = []string{
	".agents/plugins/marketplace.json",
	".claude-plugin/marketplace.json",
	".claude-plugin/plugin.json",
	".codex-plugin/plugin.json",
	"CHANGELOG.md",
	"CONTRIBUTING.md",
	"LICENSE",
	"NOTICE",
	"README.md",
	"docs/CHAT.md",
	"docs/HARNESSES.md",
	"docs/INSTALL.md",
	"docs/ztc-flow.mmd",
	"docs/ztc-flow.svg",
	"extensions/ztc.js",
	"hooks/hooks.json",
	"hooks/ztc-context.sh",
	"llm/ztc-009.llm.md",
	"package.json",
	"skills/ztc-install/SKILL.md",
	"skills/ztc-install/references/profile-template.md",
	"skills/ztc-response/SKILL.md",
	"slovnik/GUIDELINES.md",
	"slovnik/IT.md",
	"slovnik/LICENSE",
	"slovnik/README.md",
	"slovnik/RULEBOOK.md",
	"slovnik/software-development.csv",
	"slovnik/ztc-009.csv",
	"spec/ztc-009.md",
	"tools/LICENSE",
	"tools/cmd/ztcpack/main.go",
	"tools/cmd/ztcvalidate/main.go",
	"tools/go.mod",
	"tools/internal/dictionary/checker.go",
	"tools/internal/dictionary/entry.go",
	"tools/internal/dictionary/rules.go",
	"tools/internal/issue/issue.go",
	"tools/internal/spec/spec.go",
	"tools/internal/version/version.go",
}
var required = []string{".claude-plugin/plugin.json", ".codex-plugin/plugin.json", "skills/ztc-install/SKILL.md", "skills/ztc-response/SKILL.md", "llm/ztc-009.llm.md", "slovnik/ztc-009.csv"}

func collect(root, output string) ([]string, error) {
	candidates := append([]string{}, files...)
	needed := map[string]bool{}
	for _, name := range required {
		needed[filepath.FromSlash(name)] = true
	}
	var result []string
	for _, name := range candidates {
		name = filepath.FromSlash(name)
		path := filepath.Join(root, name)
		for part := name; part != "."; part = filepath.Dir(part) {
			info, err := os.Lstat(filepath.Join(root, part))
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return nil, err
			}
			if info.Mode()&os.ModeSymlink != 0 {
				return nil, fmt.Errorf("symlink is not allowed: %s", part)
			}
		}
		info, err := os.Lstat(path)
		if os.IsNotExist(err) && !needed[name] {
			continue
		}
		if err != nil {
			return nil, err
		}
		if !info.Mode().IsRegular() {
			return nil, fmt.Errorf("non-regular file is not allowed: %s", path)
		}
		if path == output {
			continue
		}
		result = append(result, name)
		delete(needed, name)
	}
	for name := range needed {
		return nil, fmt.Errorf("required file missing: %s", name)
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

// Command ztcvalidate validates the ZTČ-009 repository content.
//
// Usage: ztcvalidate -root <repo root>
// Exit codes: 0 = clean, 1 = findings, 2 = runtime error.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ztc-009/ztc-009/tools/internal/dictionary"
	"github.com/ztc-009/ztc-009/tools/internal/issue"
	"github.com/ztc-009/ztc-009/tools/internal/spec"
	"github.com/ztc-009/ztc-009/tools/internal/version"
)

func checkers(root string) ([]issue.Checker, error) {
	p := func(rel string) string { return filepath.Join(root, rel) }
	entries, err := os.ReadDir(p("slovnik"))
	if err != nil {
		return nil, err
	}
	base, err := os.Stat(p("slovnik/ztc-009.csv"))
	if err != nil {
		return nil, err
	}
	if !base.Mode().IsRegular() {
		return nil, fmt.Errorf("base dictionary is not a regular file")
	}
	checks := []issue.Checker{
		spec.Checker{Path: p("spec/ztc-009.md")},
		version.Checker{Sources: []version.Source{
			{Path: p("spec/ztc-009.md"), Pattern: regexp.MustCompile(`\|\s*\*\*Verze\*\*\s*\|\s*(\S+)\s*\|`)},
			{Path: p("llm/ztc-009.llm.md"), Pattern: regexp.MustCompile(`LLM profil v(\S+)`)},
		}},
	}
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".csv") {
			checks = append(checks, dictionary.NewChecker(p(filepath.Join("slovnik", entry.Name()))))
		}
	}
	return checks, nil
}

func main() {
	root := flag.String("root", ".", "repository root")
	localDictionary := flag.String("dictionary", "", "validate one working dictionary without repository checks")
	flag.Parse()
	var checks []issue.Checker
	if *localDictionary != "" {
		checks = []issue.Checker{dictionary.NewChecker(*localDictionary)}
	} else {
		var err error
		checks, err = checkers(*root)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}

	findings := 0
	for _, c := range checks {
		issues, err := c.Check()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", c.Name(), err)
			os.Exit(2)
		}
		for _, i := range issues {
			fmt.Println(i)
		}
		findings += len(issues)
		fmt.Fprintf(os.Stderr, "%-10s %d issue(s)\n", c.Name(), len(issues))
	}
	if findings > 0 {
		os.Exit(1)
	}
}

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

	"github.com/ztc-009/ztc-009/tools/internal/dictionary"
	"github.com/ztc-009/ztc-009/tools/internal/issue"
	"github.com/ztc-009/ztc-009/tools/internal/spec"
	"github.com/ztc-009/ztc-009/tools/internal/version"
)

func checkers(root string) []issue.Checker {
	p := func(rel string) string { return filepath.Join(root, rel) }
	return []issue.Checker{
		spec.Checker{Path: p("spec/ztc-009.md")},
		dictionary.NewChecker(p("slovnik/ztc-009.csv")),
		version.Checker{Sources: []version.Source{
			{Path: p("spec/ztc-009.md"), Pattern: regexp.MustCompile(`\|\s*\*\*Verze\*\*\s*\|\s*(\S+)\s*\|`)},
			{Path: p("llm/ztc-009.llm.md"), Pattern: regexp.MustCompile(`LLM profil v(\S+)`)},
		}},
	}
}

func main() {
	root := flag.String("root", ".", "repository root")
	flag.Parse()

	findings := 0
	for _, c := range checkers(*root) {
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

package dictionary

import (
	"os"

	"github.com/ztc-009/ztc-009/tools/internal/issue"
)

// Checker validates one dictionary file.
type Checker struct {
	Path  string
	Rules []Rule
}

// NewChecker returns a checker with the default rule set.
func NewChecker(path string) *Checker { return &Checker{Path: path, Rules: DefaultRules} }

// Name implements issue.Checker.
func (c *Checker) Name() string { return "dictionary" }

// Check implements issue.Checker.
func (c *Checker) Check() ([]issue.Issue, error) {
	f, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	entries, err := Parse(f)
	if err != nil {
		return []issue.Issue{{File: c.Path, Rule: "structure", Msg: err.Error()}}, nil
	}
	return Validate(c.Path, entries, c.Rules), nil
}

// Validate runs rules over parsed entries.
func Validate(file string, entries []Entry, rules []Rule) []issue.Issue {
	ctx := NewContext(entries)
	var out []issue.Issue
	for _, e := range entries {
		for _, r := range rules {
			for _, msg := range r.Check(e, ctx) {
				out = append(out, issue.Issue{File: file, Line: e.Line, Rule: r.ID(), Msg: msg})
			}
		}
	}
	return out
}

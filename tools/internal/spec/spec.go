// Package spec validates the rule numbering of the human-readable specification.
package spec

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/ztc-009/ztc-009/tools/internal/issue"
)

// ExpectedRules is the number of rules per section, 1:1 with ASD-STE100 Issue 9.
var ExpectedRules = map[int]int{1: 14, 2: 2, 3: 7, 4: 5, 5: 5, 6: 6, 7: 3, 8: 7, 9: 4}

var ruleLine = regexp.MustCompile(`^\*\*Pravidlo ([1-9])\.([0-9]{1,2})\*\*`)

// Rule is one rule heading found in the spec.
type Rule struct {
	Line, Section, Number int
}

// Scan returns all rule headings in document order.
func Scan(r io.Reader) ([]Rule, error) {
	var out []Rule
	sc := bufio.NewScanner(r)
	for n := 1; sc.Scan(); n++ {
		if m := ruleLine.FindStringSubmatch(sc.Text()); m != nil {
			s, _ := strconv.Atoi(m[1])
			k, _ := strconv.Atoi(m[2])
			out = append(out, Rule{Line: n, Section: s, Number: k})
		}
	}
	return out, sc.Err()
}

// Validate checks that numbering is contiguous and counts match ExpectedRules.
func Validate(file string, rules []Rule, expected map[int]int) []issue.Issue {
	var out []issue.Issue
	next := map[int]int{}
	for _, r := range rules {
		want := next[r.Section] + 1
		if r.Number != want {
			out = append(out, issue.Issue{File: file, Line: r.Line, Rule: "numbering",
				Msg: fmt.Sprintf("rule %d.%d out of order, want %d.%d", r.Section, r.Number, r.Section, want)})
		}
		next[r.Section] = r.Number
	}
	for s := 1; s <= len(expected); s++ {
		if next[s] != expected[s] {
			out = append(out, issue.Issue{File: file, Rule: "count",
				Msg: fmt.Sprintf("section %d has %d rules, want %d", s, next[s], expected[s])})
		}
	}
	return out
}

// Checker validates one spec file.
type Checker struct{ Path string }

// Name implements issue.Checker.
func (c Checker) Name() string { return "spec" }

// Check implements issue.Checker.
func (c Checker) Check() ([]issue.Issue, error) {
	f, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rules, err := Scan(f)
	if err != nil {
		return nil, err
	}
	return Validate(c.Path, rules, ExpectedRules), nil
}

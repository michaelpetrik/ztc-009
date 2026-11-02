// Package version checks that all release artifacts declare the same version.
package version

import (
	"fmt"
	"os"
	"regexp"

	"github.com/ztc-009/ztc-009/tools/internal/issue"
)

// Source is a file plus the pattern whose first group captures its version.
type Source struct {
	Path    string
	Pattern *regexp.Regexp
}

// Checker compares versions across sources.
type Checker struct{ Sources []Source }

// Name implements issue.Checker.
func (c Checker) Name() string { return "version" }

// Extract returns the version captured by pattern in data, or "".
func Extract(data []byte, p *regexp.Regexp) string {
	if m := p.FindSubmatch(data); m != nil {
		return string(m[1])
	}
	return ""
}

// Check implements issue.Checker.
func (c Checker) Check() ([]issue.Issue, error) {
	var out []issue.Issue
	ref, refPath := "", ""
	for _, s := range c.Sources {
		data, err := os.ReadFile(s.Path)
		if err != nil {
			return nil, err
		}
		v := Extract(data, s.Pattern)
		switch {
		case v == "":
			out = append(out, issue.Issue{File: s.Path, Rule: "version", Msg: "version not found"})
		case ref == "":
			ref, refPath = v, s.Path
		case v != ref:
			out = append(out, issue.Issue{File: s.Path, Rule: "version",
				Msg: fmt.Sprintf("version %s differs from %s in %s", v, ref, refPath)})
		}
	}
	return out, nil
}

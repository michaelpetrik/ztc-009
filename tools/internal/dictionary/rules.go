package dictionary

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ztc-009/ztc-009/tools/internal/spec"
)

// Rule validates a single entry against the whole dictionary context.
// New checks are added by appending a Rule to DefaultRules (open/closed).
type Rule interface {
	ID() string
	Check(e Entry, ctx *Context) []string
}

// Context holds dictionary-wide lookups shared by rules.
type Context struct {
	approvedLemmas map[string]bool
	seen           map[string]int
}

// NewContext indexes entries for cross-entry rules.
func NewContext(entries []Entry) *Context {
	c := &Context{approvedLemmas: map[string]bool{}, seen: map[string]int{}}
	for _, e := range entries {
		if e.Status == Approved {
			c.approvedLemmas[e.Lemma] = true
		}
	}
	return c
}

var (
	validPOS    = map[string]bool{"V": true, "N": true, "ADJ": true, "ADV": true, "PREP": true, "CONJ": true, "PRON": true, "NUM": true}
	validAspect = map[string]bool{"pf": true, "impf": true}
	ruleRef     = regexp.MustCompile(`^@([1-9])\.([1-9][0-9]?)$`)
)

type ruleFunc struct {
	id string
	fn func(Entry, *Context) []string
}

func (r ruleFunc) ID() string                         { return r.id }
func (r ruleFunc) Check(e Entry, c *Context) []string { return r.fn(e, c) }

// DefaultRules is the ordered rule set.
var DefaultRules = []Rule{
	ruleFunc{"required", func(e Entry, _ *Context) []string {
		var m []string
		if e.Lemma == "" {
			m = append(m, "lemma is empty")
		}
		if e.OK == "" {
			m = append(m, "ok example is empty")
		}
		return m
	}},
	ruleFunc{"pos", func(e Entry, _ *Context) []string {
		if !validPOS[e.POS] {
			return []string{fmt.Sprintf("invalid pos %q", e.POS)}
		}
		return nil
	}},
	ruleFunc{"aspect", func(e Entry, _ *Context) []string {
		switch {
		case e.POS == "V" && !validAspect[e.Aspect]:
			return []string{fmt.Sprintf("verb needs vid pf|impf, got %q", e.Aspect)}
		case e.POS != "V" && e.Aspect != "":
			return []string{"vid is allowed only for verbs"}
		}
		return nil
	}},
	ruleFunc{"status", func(e Entry, _ *Context) []string {
		switch e.Status {
		case Approved:
			var m []string
			if e.Meaning == "" {
				m = append(m, "approved entry needs meaning")
			}
			if len(e.ReplaceWith) > 0 {
				m = append(m, "approved entry must not have replace_with")
			}
			return m
		case Unapproved:
			var m []string
			if len(e.ReplaceWith) == 0 {
				m = append(m, "unapproved entry needs replace_with")
			}
			if e.Bad == "" {
				m = append(m, "unapproved entry needs bad example")
			}
			if e.Meaning != "" {
				m = append(m, "unapproved entry must not have meaning")
			}
			return m
		default:
			return []string{fmt.Sprintf("invalid status %q", e.Status)}
		}
	}},
	ruleFunc{"replacement", func(e Entry, c *Context) []string {
		var m []string
		for _, r := range e.ReplaceWith {
			validRule := false
			if match := ruleRef.FindStringSubmatch(r); match != nil {
				section, _ := strconv.Atoi(match[1])
				number, _ := strconv.Atoi(match[2])
				validRule = number <= spec.ExpectedRules[section]
			}
			if !validRule && !c.approvedLemmas[r] {
				m = append(m, fmt.Sprintf("replacement %q is not an approved lemma or @X.Y rule ref", r))
			}
		}
		return m
	}},
	ruleFunc{"unique", func(e Entry, c *Context) []string {
		k := e.Key()
		if prev, ok := c.seen[k]; ok {
			return []string{fmt.Sprintf("duplicate entry %s (first at line %d)", k, prev)}
		}
		c.seen[k] = e.Line
		return nil
	}},
}

// Package issue defines the shared validation finding type.
package issue

import "fmt"

// Issue is one validation finding.
type Issue struct {
	File string
	Line int // 0 = whole file
	Rule string
	Msg  string
}

func (i Issue) String() string {
	if i.Line > 0 {
		return fmt.Sprintf("%s:%d: [%s] %s", i.File, i.Line, i.Rule, i.Msg)
	}
	return fmt.Sprintf("%s: [%s] %s", i.File, i.Rule, i.Msg)
}

// Checker produces issues. Every validator implements it.
type Checker interface {
	Name() string
	Check() ([]Issue, error)
}

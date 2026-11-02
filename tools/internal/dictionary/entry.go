// Package dictionary parses and validates the ZTČ-009 dictionary CSV.
package dictionary

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// Header is the required column order.
var Header = []string{"lemma", "pos", "vid", "status", "meaning", "replace_with", "en_ste", "ok", "bad"}

// Status values.
const (
	Approved   = "approved"
	Unapproved = "unapproved"
)

// Entry is one dictionary row. Line is the 1-based CSV line number.
type Entry struct {
	Line        int
	Lemma       string
	POS         string
	Aspect      string
	Status      string
	Meaning     string
	ReplaceWith []string
	EnSTE       string
	OK          string
	Bad         string
}

// Key identifies an entry uniquely.
func (e Entry) Key() string { return e.Lemma + "/" + e.POS + "/" + e.Aspect }

// Parse reads the CSV. It fails only on structural errors (header, column count).
func Parse(r io.Reader) ([]Entry, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = len(Header)
	head, err := cr.Read()
	if err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}
	if strings.Join(head, ",") != strings.Join(Header, ",") {
		return nil, fmt.Errorf("header mismatch: got %q, want %q", strings.Join(head, ","), strings.Join(Header, ","))
	}
	var out []Entry
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		line, _ := cr.FieldPos(0)
		out = append(out, Entry{
			Line:        line,
			Lemma:       strings.TrimSpace(rec[0]),
			POS:         strings.TrimSpace(rec[1]),
			Aspect:      strings.TrimSpace(rec[2]),
			Status:      strings.TrimSpace(rec[3]),
			Meaning:     strings.TrimSpace(rec[4]),
			ReplaceWith: splitList(rec[5]),
			EnSTE:       strings.TrimSpace(rec[6]),
			OK:          strings.TrimSpace(rec[7]),
			Bad:         strings.TrimSpace(rec[8]),
		})
	}
}

func splitList(s string) []string {
	var out []string
	for _, p := range strings.Split(s, "|") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

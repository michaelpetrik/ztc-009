package dictionary

import (
	"strings"
	"testing"
)

const head = "lemma,pos,vid,status,meaning,replace_with,en_ste,ok,bad\n"

func validate(t *testing.T, body string) []string {
	t.Helper()
	es, err := Parse(strings.NewReader(head + body))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var rules []string
	for _, i := range Validate("x.csv", es, DefaultRules) {
		rules = append(rules, i.Rule)
	}
	return rules
}

func TestParseHeaderMismatch(t *testing.T) {
	if _, err := Parse(strings.NewReader("a,b\n")); err == nil {
		t.Fatal("expected error")
	}
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name, body string
		want       []string
	}{
		{"valid approved", "zavřít,V,pf,approved,m,,close,ok,\n", nil},
		{"valid unapproved", "zavřít,V,pf,approved,m,,,ok,\nuzavřít,V,pf,unapproved,,zavřít,,ok,bad\n", nil},
		{"rule ref replacement", "provést,V,pf,unapproved,,@3.7,,ok,bad\n", nil},
		{"last section one rule", "x,N,,unapproved,,@1.14,,ok,bad\n", nil},
		{"last section nine rule", "x,N,,unapproved,,@9.4,,ok,bad\n", nil},
		{"nonexistent rule", "x,N,,unapproved,,@9.99,,ok,bad\n", []string{"replacement"}},
		{"zero rule", "x,N,,unapproved,,@1.0,,ok,bad\n", []string{"replacement"}},
		{"bad pos", "x,Q,,approved,m,,,ok,\n", []string{"pos"}},
		{"verb without aspect", "x,V,,approved,m,,,ok,\n", []string{"aspect"}},
		{"aspect on noun", "x,N,pf,approved,m,,,ok,\n", []string{"aspect"}},
		{"approved no meaning", "x,N,,approved,,,,ok,\n", []string{"status"}},
		{"unapproved no replacement", "x,N,,unapproved,,,,ok,bad\n", []string{"status"}},
		{"unknown replacement", "x,N,,unapproved,,nic,,ok,bad\n", []string{"replacement"}},
		{"duplicate", "x,N,,approved,m,,,ok,\nx,N,,approved,m,,,ok,\n", []string{"unique"}},
		{"missing ok", "x,N,,approved,m,,,,\n", []string{"required"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := validate(t, c.body)
			if strings.Join(got, ",") != strings.Join(c.want, ",") {
				t.Fatalf("got %v, want %v", got, c.want)
			}
		})
	}
}

package spec

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	exp := map[int]int{1: 2, 2: 1}
	cases := []struct {
		name, doc string
		want      int
	}{
		{"ok", "**Pravidlo 1.1** a\n**Pravidlo 1.2** b\n**Pravidlo 2.1** c\n", 0},
		{"gap", "**Pravidlo 1.1** a\n**Pravidlo 1.3** b\n**Pravidlo 2.1** c\n", 2},
		{"missing section", "**Pravidlo 1.1** a\n**Pravidlo 1.2** b\n", 1},
		{"ignores body text", "text Pravidlo 1.1\n**Pravidlo 1.1** a\n**Pravidlo 1.2** b\n**Pravidlo 2.1** c\n", 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rules, err := Scan(strings.NewReader(c.doc))
			if err != nil {
				t.Fatal(err)
			}
			if got := Validate("s.md", rules, exp); len(got) != c.want {
				t.Fatalf("got %d issues %v, want %d", len(got), got, c.want)
			}
		})
	}
}

package main

import (
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	for _, tc := range []struct {
		name, text       string
		trivial, invalid bool
	}{
		{"valid", "CI (průběžná integrace) kontroluje změny. Zkontrolujte výsledek.", false, false},
		{"bold acronym", "**CI** (průběžná integrace) běží.", false, false},
		{"decimal version", "Verze 1.2.3 používá hodnotu 1,5.", false, false},
		{"heading punctuation", "# Stav; výsledek", false, true},
		{"heading acronym", "# CI", false, true},
		{"heading length", "# " + strings.Repeat("slovo ", 30), false, true},
		{"semicolon", "Systém běží; chyba trvá.", false, true},
		{"abbreviation", "Zkontrolujte např. soubor.", false, true},
		{"short abbreviation", "To je tj. chyba.", false, true},
		{"unexpanded", "CI kontroluje změny.", false, true},
		{"wrong expansion", "VM (databáze) běží.", false, true},
		{"second use", "VM (virtuální stroj) běží. Zastavte VM.", false, false},
		{"long", strings.Repeat("slovo ", 21) + ".", false, true},
		{"paragraph", strings.Repeat("Systém běží. ", 7), false, true},
		{"paragraph breaks", strings.Repeat("Systém běží.\n\n", 7), false, false},
		{"code", "Spusťte `a;b`.\n```js\nCI; např.;\n```", false, true},
		{"quote", "Citace: „CI; např.“", false, true},
		{"link", "Zkontrolujte [soubor](https://example.com/a;b).", false, false},
		{"link label", "[CI](https://example.com) běží.", false, true},
		{"unclosed fence", "```js\nSystém běží;", false, true},
		{"nested fence", "````text\n```js\nCI;\n```\n````", false, true},
		{"trivial", "Ano, CI běží.", true, false},
		{"trivial multi", "Ano. Ale chyba trvá.", true, true},
		{"trivial list", "Ano.\nNe.", true, true},
		{"trivial empty", "", true, true},
		{"fake heading", "#" + strings.Repeat("slovo ", 40), false, true},
		{"long parentheses", "Stav (" + strings.Repeat("slovo ", 80) + ").", false, true},
		{"invisible abbreviation", "Zkontrolujte na\u200bpř. soubor.", false, true},
		{"blockquote bypass", "> CI; např. chyba.", false, true},
		{"passive construction", "Ventil musí být uzavřen.", false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			issues := check(tc.text, tc.trivial)
			if (len(issues) > 0) != tc.invalid {
				t.Fatalf("findings=%v, invalid=%v", issues, tc.invalid)
			}
		})
	}
}

func TestSourceProtection(t *testing.T) {
	source := "Citace: „CI; např. chyba.“\n```js\nconsole.log('CI;');\n```"
	_, spans, err := prose(source, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	trusted := map[string]bool{}
	for _, span := range spans {
		trusted[span] = true
	}
	if issues := checkProtected(source, false, trusted, false); len(issues) != 0 {
		t.Fatal(issues)
	}
	if issues := checkProtected("„CI; jiná chyba.“", false, trusted, false); len(issues) == 0 {
		t.Fatal("altered quote bypassed")
	}
	if issues := checkProtected("```js\nconst x = 1;\n```", false, nil, true); len(issues) != 0 {
		t.Fatal(issues)
	}
	if issues := checkProtected("CI; chyba.\n```js\nconst x = 1;\n```", false, nil, true); len(issues) == 0 {
		t.Fatal("code permission exempted prose")
	}
}

package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var inlineCode = regexp.MustCompile("`+[^`\n]+`+")
var links = regexp.MustCompile(`\[([^\]]+)\]\([^\s)]+\)`)
var urls = regexp.MustCompile(`https?://[^\s]+`)
var quotes = regexp.MustCompile(`„[^“]*“|"[^"\n]*"`)
var parentheses = regexp.MustCompile(`\([^()]*\)`)
var listPrefix = regexp.MustCompile(`^\s*(?:[-*+] |\d+[.)] )`)
var abbreviations = regexp.MustCompile(`(?i)(^|[^\pL\pN_])(?:tj\.|tzn\.|např\.|atd\.|apod\.|resp\.|cca|tzv\.|max\.|min\.)([^\pL\pN_]|$)`)
var decimals = regexp.MustCompile(`[0-9]+(?:[.,][0-9]+)+`)

func prose(text string) ([]string, error) {
	var lines []string
	var fence byte
	var width int
	for _, line := range strings.Split(text, "\n") {
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "```") || strings.HasPrefix(s, "~~~") {
			c := s[0]
			n := len(s) - len(strings.TrimLeft(s, string(c)))
			if fence == 0 {
				fence, width = c, n
				lines = append(lines, "")
			} else if c == fence && n >= width && strings.TrimSpace(s[n:]) == "" {
				fence = 0
			}
			continue
		}
		if fence != 0 || strings.HasPrefix(s, ">") {
			continue
		}
		s = inlineCode.ReplaceAllString(s, " KÓD ")
		s = links.ReplaceAllString(s, "$1")
		s = urls.ReplaceAllString(s, " ODKAZ ")
		s = quotes.ReplaceAllString(s, " CITACE ")
		if listPrefix.MatchString(s) {
			lines = append(lines, "")
			s = listPrefix.ReplaceAllString(s, "")
		}
		lines = append(lines, s)
	}
	if fence != 0 {
		return nil, fmt.Errorf("neuzavřený blok kódu")
	}
	return lines, nil
}

func sentences(text string) []string {
	text = decimals.ReplaceAllString(text, "ČÍSLO")
	text = parentheses.ReplaceAllString(text, " VYSVĚTLENÍ ")
	return strings.FieldsFunc(text, func(r rune) bool { return r == '.' || r == '!' || r == '?' || r == ':' })
}

func wordCount(s string) int {
	n := 0
	for _, word := range strings.Fields(s) {
		if strings.IndexFunc(word, func(r rune) bool { return unicode.IsLetter(r) || unicode.IsNumber(r) }) >= 0 {
			n++
		}
	}
	return n
}

func check(text string, trivial bool) []string {
	lines, err := prose(text)
	if err != nil {
		return []string{err.Error()}
	}
	clean := strings.Join(lines, "\n")
	countLines := append([]string{}, lines...)
	for i, line := range countLines {
		if strings.HasPrefix(line, "#") {
			countLines[i] = ""
		}
	}
	counted := strings.Join(countLines, "\n")
	parts := sentences(counted)
	if trivial {
		if len(parts) == 1 && wordCount(parts[0]) > 0 && wordCount(parts[0]) <= 25 && !strings.Contains(text, "\n") {
			return nil
		}
		return []string{"Výjimka trivial platí jen pro jednu neprázdnou větu na jednom řádku, nejvýše 25 slov."}
	}
	var findings []string
	if strings.Contains(clean, ";") {
		findings = append(findings, "8.1: próza obsahuje středník")
	}
	if abbreviations.MatchString(clean) {
		findings = append(findings, "D-6: próza obsahuje zakázanou obecnou zkratku")
	}
	for _, sentence := range parts {
		if n := wordCount(sentence); n > 20 {
			findings = append(findings, fmt.Sprintf("5.1/6.3: věta má %d slov; kontrolní režim vyžaduje nejvýše 20", n))
			break
		}
	}
	for _, paragraph := range strings.Split(counted, "\n\n") {
		n := 0
		for _, sentence := range sentences(paragraph) {
			if wordCount(sentence) > 0 {
				n++
			}
		}
		if n > 6 {
			findings = append(findings, "6.6: odstavec má více než 6 vět")
			break
		}
	}
	for _, term := range [][2]string{
		{"CI", "průběžná integrace"}, {"VM", "virtuální stroj"}, {"FE", "frontend"}, {"BE", "backend"},
		{"UI", "uživatelské rozhraní"}, {"DB", "databáze"}, {"PR", "požadavek na sloučení"},
	} {
		acronym, expansion := term[0], term[1]
		re := regexp.MustCompile(`\b` + acronym + `\b`)
		if pos := re.FindStringIndex(clean); pos != nil {
			rest := strings.TrimSpace(strings.TrimLeft(clean[pos[1]:], "*_"))
			if !strings.HasPrefix(rest, "("+expansion+")") {
				findings = append(findings, "10.4: první "+acronym+" musí mít tvar "+acronym+" ("+expansion+")")
			}
		}
	}
	return findings
}

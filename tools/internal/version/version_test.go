package version

import (
	"regexp"
	"testing"
)

func TestExtract(t *testing.T) {
	p := regexp.MustCompile(`profil v(\S+)`)
	if got := Extract([]byte("ZTČ-009 LLM profil v1.0.0-rc1\n"), p); got != "1.0.0-rc1" {
		t.Fatalf("got %q", got)
	}
	if got := Extract([]byte("nic"), p); got != "" {
		t.Fatalf("got %q", got)
	}
}

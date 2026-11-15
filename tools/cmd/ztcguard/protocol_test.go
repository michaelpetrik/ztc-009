package main

import "testing"

func TestOutputSchema(t *testing.T) {
	for _, input := range []string{"", "null", "[]", "garbage", "{}{}", `{"continue":true}`, `{"decision":"allow","reason":"ok"}`, `{"hookSpecificOutput":{}}`} {
		if _, err := validateOutput([]byte(input)); err == nil {
			t.Errorf("accepted %q", input)
		}
	}
	for _, input := range []string{`{}`, `{"decision":"block","reason":"fix"}`, `{"continue":false,"stopReason":"failed","systemMessage":"failed"}`, `{"hookSpecificOutput":{"hookEventName":"UserPromptSubmit","additionalContext":"context"}}`} {
		if _, err := validateOutput([]byte(input)); err != nil {
			t.Errorf("rejected %q: %v", input, err)
		}
	}
}

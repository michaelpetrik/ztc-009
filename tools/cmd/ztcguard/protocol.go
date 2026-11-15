package main

import (
	"encoding/json"
	"errors"
)

func validateOutput(data []byte) ([]byte, error) {
	var value map[string]json.RawMessage
	if len(data) > 65536 || json.Unmarshal(data, &value) != nil || value == nil {
		return nil, errors.New("invalid guard JSON")
	}
	valid := len(value) == 0
	if len(value) == 2 {
		var decision, reason string
		if json.Unmarshal(value["decision"], &decision) == nil && json.Unmarshal(value["reason"], &reason) == nil {
			valid = decision == "block" && reason != ""
		}
	}
	if len(value) == 3 {
		var cont *bool
		var reason, message string
		if json.Unmarshal(value["continue"], &cont) == nil && json.Unmarshal(value["stopReason"], &reason) == nil && json.Unmarshal(value["systemMessage"], &message) == nil {
			valid = cont != nil && !*cont && reason != "" && message != ""
		}
	}
	if len(value) == 1 {
		var context map[string]string
		if json.Unmarshal(value["hookSpecificOutput"], &context) == nil {
			valid = len(context) == 2 && context["hookEventName"] == "UserPromptSubmit" && context["additionalContext"] != ""
		}
	}
	if !valid {
		return nil, errors.New("invalid guard decision schema")
	}
	return json.Marshal(value)
}

package leetcode

import "testing"

func TestValidParentheses(t *testing.T) {
	s := "{[]}"
	valid := ValidParentheses(s)
	if !valid {
		t.Fatalf("%s should be valid", s)
	}

	s = "{}[]()"
	valid = ValidParentheses(s)
	if !valid {
		t.Fatalf("%s should be valid", s)
	}

	s = "{[()]}"
	valid = ValidParentheses(s)
	if !valid {
		t.Fatalf("%s should be valid", s)
	}

	s = "{[(])}"
	valid = ValidParentheses(s)
	if valid {
		t.Fatalf("%s should be invalid", s)
	}

}

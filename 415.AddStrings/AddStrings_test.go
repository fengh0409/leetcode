package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	num1, num2 := "11", "123"
	expected := "134"
	result := AddStrings(num1, num2)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}

	num1, num2 = "456", "77"
	expected = "533"
	result = AddStrings(num1, num2)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}

	num1, num2 = "0", "0"
	expected = "0"
	result = AddStrings(num1, num2)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}
}

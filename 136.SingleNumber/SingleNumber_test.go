package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	expected := 1
	result := SingleNumber(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{4, 1, 2, 1, 2}
	expected = 4
	result = SingleNumber(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}
}

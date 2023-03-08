package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	result := MaxSubArray(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{1}
	expected = 1
	result = MaxSubArray(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{5, 4, -1, 7, 8}
	expected = 23
	result = MaxSubArray(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{5, -6, 4, -1, 7, 8}
	expected = 18
	result = MaxSubArray(nums)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}
}

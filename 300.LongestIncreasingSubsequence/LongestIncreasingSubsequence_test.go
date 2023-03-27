package leetcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4
	result := LengthOfLIS(nums)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}

	nums = []int{0, 1, 0, 3, 2, 3}
	expected = 4
	result = LengthOfLIS(nums)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}

	nums = []int{7, 7, 7, 7, 7, 7, 7}
	expected = 1
	result = LengthOfLIS(nums)
	if result != expected {
		t.Fatalf("should be %v,but is %v", expected, result)
	}
}

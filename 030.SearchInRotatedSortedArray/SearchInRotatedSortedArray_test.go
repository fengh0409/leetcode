package leetcode

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4
	result := Search(nums, target)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	expected = -1
	result = Search(nums, target)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{1}
	target = 0
	expected = -1
	result = Search(nums, target)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}
}

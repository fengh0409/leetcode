package leetcode

import "testing"

func TestMedianOfTwoSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var expected float64 = 4
	result := MedianOfTwoSortedArrays(nums1, nums2)
	if result != expected {
		t.Fatalf("should be %v, but is %v", expected, result)
	}

	nums1 = []int{1, 9}
	nums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected = 5
	result = MedianOfTwoSortedArrays(nums1, nums2)
	if result != expected {
		t.Fatalf("should be %v, but is %v", expected, result)
	}
}

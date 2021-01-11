package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	var pos = TwoSum(nums, target)
	if len(pos) != 2 || nums[pos[0]]+nums[pos[1]] != target {
		t.Fatalf("result: %v is wrong", pos)
	}
}

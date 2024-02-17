package leetcode

import (
	"math"
	"sort"
)

// ThreeSumClosest 三数之和
func ThreeSumClosest(nums []int, target int) int {
	// 先排序
	sort.Ints(nums)

	var result = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		// 双指针
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

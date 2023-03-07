package leetcode

// FindKthLarget 数组中的第K个最大元素
func FindKthLargest(nums []int, k int) int {
	target := len(nums) - k
	left, right := 0, len(nums)-1

	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[target]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

// 基于快排
func partition(nums []int, i, j int) int {
	pivot := i
	left, right := pivot+1, j
	for k := left; k <= right; k++ {
		if nums[k] < nums[pivot] {
			nums[k], nums[left] = nums[left], nums[k]
			left++
		}
	}
	nums[pivot], nums[left-1] = nums[left-1], nums[pivot]

	return left - 1
}

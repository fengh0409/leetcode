package leetcode

// MergeSort 归并排序，时间复杂度O(nlogn)，稳定
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 递归地分而治之
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	l, r := 0, 0
	result := make([]int, 0, l+r)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	// 归并剩余元素
	if l < len(left) {
		result = append(result, left[l:]...)
	}
	if r < len(right) {
		result = append(result, right[r:]...)
	}

	return result
}

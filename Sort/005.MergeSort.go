package leetcode

import "fmt"

// MergeSort 归并排序，时间复杂度O(nlogn)，稳定
func MergeSort(nums []int) []int {
	fmt.Println(nums)
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
	fmt.Println("merge before===>", left, right)
	l, r := 0, 0
	result := make([]int, 0, l+r)
	// 最后一次归并变成合并两个有序数组
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

	fmt.Println("merge after===>", result)
	return result
}

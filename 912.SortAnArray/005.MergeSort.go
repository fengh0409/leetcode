package leetcode

import "fmt"

// MergeSort 归并排序，时间复杂度O(nlogn)，稳定
func MergeSort(nums []int) []int {
	fmt.Println(nums)
	// 递归退出条件
	if len(nums) <= 1 {
		return nums
	}

	// 递归地分而治之
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}

// 最后一次归并变成合并两个有序数组
func merge(left, right []int) []int {
	fmt.Println("merge before===>", left, right)
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}

	fmt.Println("merge after===>", result)
	return result
}

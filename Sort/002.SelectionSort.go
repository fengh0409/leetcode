package leetcode

// SelectionSort 选择排序，时间复杂度O(n²)，不稳定
func SelectionSort(nums []int) []int {
	// 每次遍历，将第一个数和后面的元素进行比较，保存最小元素的索引
	// 遍历完成后，将第一个数和最小元素位置交换
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}

	// 附：每次比较完都交换
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[j] < nums[min] {
	//			nums[i], nums[j] = nums[j], nums[i]
	//		}
	//	}
	//}

	return nums
}

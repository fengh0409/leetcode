package leetcode

// QuickSort 快速排序，时间复杂度O(nlogn)，最坏情况O(n^2)，不稳定
func QuickSort(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		// 取第一个值作为基准值
		pivot := left
		// 基准值后面的第一个数
		index := pivot + 1
		// 遍历基准值后面所有的数据，和基准值进行比较
		for i := index; i <= right; i++ {
			if nums[i] < nums[pivot] {
				nums[index], nums[i] = nums[i], nums[index]
				index++
			}
		}
		// 比较一趟下来，index前半部分都比基准值小，index后半部分都比基准值大
		nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
		// 重新设置pivot，进行分治
		mid := index - 1
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}

	return nums
}

package leetcode

// QuickSort 快速排序，时间复杂度O(nlogn)，最坏情况O(n^2)，不稳定，但相比其他排序更快
func QuickSort(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

// 基本快排
func quickSort(nums []int, i, j int) []int {
	if i < j {
		// 取第一个值作为基准值
		pivot := i
		// 基准值后面的第一个数
		left, right := pivot+1, j
		// 遍历基准值后面所有的数据，和基准值进行比较，比基准值小的往左边放
		for k := left; k <= right; k++ {
			if nums[k] < nums[pivot] {
				nums[left], nums[k] = nums[k], nums[left]
				left++
			}
		}
		// 比较一趟下来，left前半部分都比基准值小，left后半部分都比基准值大
		// 交换基准值与分治位置
		nums[pivot], nums[left-1] = nums[left-1], nums[pivot]
		// 重新设置边界，进行分治
		mid := left - 1

		quickSort(nums, i, mid-1)
		quickSort(nums, mid+1, j)
	}

	return nums
}

// 解法二：双指针
//func quickSort(nums []int, i, j int) []int {
//	if i < j {
//		pivot := i
//		left, right := i+1, j
//		for {
//			for left <= right && nums[left] < nums[pivot] {
//				left++
//			}
//			for left <= right && nums[right] > nums[pivot] {
//				right--
//			}
//
//			if left >= right {
//				break
//			}
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//			right--
//		}
//
//		nums[pivot], nums[left-1] = nums[left-1], nums[pivot]
//		mid := left - 1
//		quickSort(nums, i, mid-1)
//		quickSort(nums, mid+1, j)
//	}
//
//	return nums
//}

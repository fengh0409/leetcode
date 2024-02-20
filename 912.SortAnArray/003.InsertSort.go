package leetcode

// InsertSort 插入排序，时间复杂度O(n²)，稳定
func InsertSort(nums []int) []int {
	// 假设前面n个元素已排好序，后面未排序的元素i和前面已排序元素i依次从后往前进行比较，
	// 若元素i小于元素j，则将元素j往后移一位，然后继续比较，
	// 直到元素i大于元素j，则将元素i插入在元素j的后面
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1
		for j >= 0 {
			if cur < nums[j] {
				nums[j+1] = nums[j]
				j--
			} else {
				break
			}
		}
		nums[j+1] = cur
	}

	return nums
}

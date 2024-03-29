package leetcode

// RemoveDuplicatesFromSortArray 删除排序数组中的重复项
func RemoveDuplicatesFromSortArray(nums []int) int {
	// 双指针法，时间复杂度O(n)，空间复杂度O(1)
	// 注意数组是有序的，也不需要考虑后面多余的元素
	// 这种方式其实就是用后面新的元素覆盖前面重复的元素
	var i = 0
	for j := 1; j < len(nums); j++ {
		// 每次遍历判断与之前的元素是否重复，若不重复，说明是新的元素，
		// 然后用后面的元素覆盖前面的，并将两个数的下标向后移一位
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}

	// 最后的结果需要+1
	return i + 1

	//var m = make(map[int]int)
	//for i := len(nums) - 1; i >= 0; i-- {
	//	if k, ok := m[nums[i]]; ok {
	//		copy(nums[i:], nums[k:])
	//		nums = nums[0 : len(nums)-1]
	//	}
	//	m = map[int]int{}
	//	m[nums[i]] = i
	//}

	//return len(nums)
}

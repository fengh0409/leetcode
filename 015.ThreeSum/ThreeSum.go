package leetcode

import "sort"

// ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	// 1. 对数组进行排序。
	// 2. 遍历排序后数组：
	//    2.1 若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
	//    2.2 对于重复元素：跳过，避免出现重复解
	//    2.3 令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
	//        2.3.1 当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。
	//              并同时将 L,R 移到下一位置，寻找新的解
	//        2.3.2 若和大于 0，说明 nums[R] 太大，R 左移
	//        2.3.3 若和小于 0，说明 nums[L] 太小，L 右移

	// 先排序
	sort.Ints(nums)

	// 双指针
	var res = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		// 若第一个数大于0，则后面的数一定大于0，其和不可能为0，因此跳出循环
		if nums[i] > 0 {
			break
		}
		// 从第二次遍历开始，如果这个数与上一次遍历的相同，则表示重复数据，需要跳过，以去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 左指针，表示第二个数
		left := i + 1
		// 右指针，表示第三个数
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 找到三数之和为0的数后，左右指针均需要往中间靠拢一位
				left++
				right--
				// 判断移动后的左指针的数据与上一次是否相等，相等表示重复了，需要继续向右移动指针
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 判断移动后的右指针的数据与上一次是否相等，相等表示重复了，需要继续向左移动指针
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[i]+nums[left]+nums[right] > 0 {
				// 如果三数之和大于0，则需要第三个数往左移动，因为数组已经排好序，右边的数较大，
				// 只能将右指针往左移动，才可能使三数之和为0
				right--
			} else {
				// 如果三数之和小于0，则需要第二个数往右移动，因为数组已经排好序，左边数较小
				// 只能将左指针往右移动，才可能使三数之和为0
				left++
			}
		}
	}
	return res
}

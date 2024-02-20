package leetcode

// BubbleSort 冒泡排序，时间复杂度O(n²)，稳定
// 所谓稳定性，即如果a原本在b前面，而a=b，排序之后a仍然在b的前面
func BubbleSort(nums []int) []int {
	// 比较相邻的两个元素，如果第一个比第二个大，就交换顺序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

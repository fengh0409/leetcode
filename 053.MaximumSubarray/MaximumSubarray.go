package leetcode

// MaxSubArray 最大子数组和
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果累加后比原来的数大，说明新加的数都是正值，需要保留
		// 但是如果前面累加后还不如自身大，则需要把前面的都丢掉，重新计算
		// 例如：[2,3,-1,4,3]
		// 2+3=5,5>3,保留
		// 5+(-1)=4,4<5,丢弃，此时nums[i]仍然为-1，但最大值仍然为5
		// 4+(-1)=3,3>-1,保留
		// 3+3=6,6>3,保留，最大和为6
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

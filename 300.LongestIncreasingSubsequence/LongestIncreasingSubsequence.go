package leetcode

// LengthOfLIS 最长递增子序列
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxLen int
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 每个dp元素设为1
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

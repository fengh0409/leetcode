package leetcode

// ContainerWithMostWater 盛最多水的容器
func ContainerWithMostWater(height []int) int {
	// 双指针法：数组中的每个元素代表板的高度，能盛多少水取决于短板，
	// 因此，每次移动头指针或尾指针，要看谁的板高度最短，
	// 头指针短则向右移动指针，尾指针短则向左移动指针，
	// 两个指针向中间靠拢，当两个指针交汇时则结束。
	var maxArea int
	var left = 0
	var right = len(height) - 1
	for left < right {
		var area int
		if height[left] > height[right] {
			area = (right - left) * height[right]
			right--
		} else {
			area = (right - left) * height[left]
			left++
		}
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

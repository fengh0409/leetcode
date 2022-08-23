package leetcode

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	// 可以通过遍历，时间复杂度是O(n²)，但不是最佳的解法
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			pos[0], pos[1] = i, j
	//          break
	//		}
	//	}
	//}

	// 使用map，每次遍历，到map中寻找v1的另一半v2是否存在
	// 若不存在则把v1写入map，等到下次遍历到v2时，就可以在map中找到另一半v1了
	// 时间复杂度是O(n)
	var m = make(map[int]int)
	for i, v1 := range nums {
		v2 := target - v1
		if j, ok := m[v2]; ok {
			return []int{i, j}
		}
		m[v1] = i
	}

	return nil
}

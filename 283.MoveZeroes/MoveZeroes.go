package leetcode

func MoveZeroes(nums []int) {
	insertPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertPos] = nums[i]
			insertPos++
		}
	}

	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

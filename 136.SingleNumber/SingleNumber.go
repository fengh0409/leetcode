package leetcode

// SingleNumber 只出现一次的数字
func SingleNumber(nums []int) int {
	// 异或运算：一个数与自身异或运算结果总是0，一个数与0异或元素结果等于其本身
	// 由于数组中只有一个数字出现一次，其他数字均出现两次，那把数组中的所有元素进行异或运算后，结果就是只出现了一次的数值
	var val = 0
	for _, num := range nums {
		val ^= num
	}
	return val
}

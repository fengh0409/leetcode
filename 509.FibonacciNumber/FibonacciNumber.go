package leetcode

//var m = map[int]int{}

func Fibonacci(n int) int {
	// 自底向上遍历，时间复杂度O(n)
	if n < 2 {
		return n
	}

	var f0, f1 = 0, 1
	for i := 1; i < n; i++ {
		// 交换变量，无需借助辅助变量
		f0, f1 = f1, f0+f1
	}

	return f1

	// 1.递归，自顶向下，二叉树结构，大量的重复计算，时间复杂度2的n次方
	//if n < 2 {
	//	return n
	//}

	//return Fibonacci(n-1) + Fibonacci(n-2)

	// 2.递归+备忘录算法，存储每次计算结果，避免重复计算，时间复杂度O(n)
	//if n < 2 {
	//	m[n] = n
	//	return n
	//}
	//if v, ok := m[n]; ok {
	//	return v
	//}

	//value := Fibonacci(n-1) + Fibonacci(n-2)
	//m[n] = value
	//return value

}

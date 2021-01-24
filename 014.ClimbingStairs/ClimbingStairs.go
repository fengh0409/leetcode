package leetcode

//var m = map[int]int{}

// ClimbingStairs 爬楼梯
func ClimbingStairs(n int) int {
	// 动态规划问题
	// 分解问题：假设只差一步就能到达第n阶楼顶，那么可以分为两种情况，一是处于第n-1阶，再爬1个台阶即可到达，二是处于第n-2阶，再爬2个台阶即可到达
	// 按照第一种情况，假设到达第n-1阶，有x种走法，那么到达第n阶也是x种走法
	// 按照第二种情况，假设到达第n-2阶，有y种走法，那么到达第n阶也是y种走法
	// 因此，第n阶与第n-1阶和第n-2阶的关系是，到达第n阶的走法为第n-1阶与第n-2阶的走法之和
	// 用函数表达式来表示，也就是F(n)=F(n-1)+F(n-2)，同理到达第n-1阶的走法F(n-1)=F(n-2)+F(n-3)，n>=3
	// 也就是说F(n)的结果，需要从F(n-1)和F(n-2)转移过来，F(n)=F(n-1)+F(n-2)即为状态转移方程
	// F(n)的最优子结构为F(n-1)和F(n-2)，当只有1阶或2阶台阶时，F(1)=1，F(2)=2，这也是问题的边界
	// 根据状态转移方程，首先看出可以使用递归求解

	// 递归，时间复杂度2的n次方
	//if n < 3 {
	//	return n
	//}

	//return ClimbingStairs(n-1) + ClimbingStairs(n-2)

	// 但是递归求解的时间复杂度为2的n次方，效率太低，因为其中做了很多重复计算，可以做如下的优化
	// 备忘录算法，把结果存起来，时间复杂度和空间复杂度均为O(n)
	//if n < 3 {
	//	return n
	//}
	//if v, ok := m[n]; ok {
	//	return v
	//}

	//value := ClimbingStairs(n-1) + ClimbingStairs(n-2)
	//m[n] = value

	//return value

	// 再次优化，自底向上遍历，时间复杂度O(n)，空间复杂度O(1)
	if n < 3 {
		return n
	}

	var f1, f2 = 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

package leetcode

// MergeSortedArray 合并两个有序数组
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// 反向双指针法，从尾部开始算，先将最大的数放在最后
	p1, p2 := m-1, n-1
	for t := m + n - 1; p1 >= 0 || p2 >= 0; t-- {
		if p1 == -1 {
			nums1[t] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[t] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[t] = nums1[p1]
			p1--
		} else {
			nums1[t] = nums2[p2]
			p2--
		}
	}

	//// 双指针法
	//var sorted = make([]int, 0, m+n)
	//p1, p2 := 0, 0
	//for {
	//	// p1与m相等时，表示nums1数组遍历完了，此时把nums2剩下的元素append即可
	//	if p1 == m {
	//		sorted = append(sorted, nums2[p2:]...)
	//		break
	//	}
	//	// p2与n相等时，表示nums2数组遍历完了，此时把nums1剩下的元素append即可
	//	if p2 == n {
	//		sorted = append(sorted, nums1[p1:]...)
	//		break
	//	}
	//	if nums1[p1] < nums2[p2] {
	//		sorted = append(sorted, nums1[p1])
	//		p1++
	//	} else {
	//		sorted = append(sorted, nums2[p2])
	//		p2++
	//	}
	//}
	//copy(nums1, sorted)

	// 解法一：合并后直接排序，但时间复杂度较高
	//for i := 0; i < n; i++ {
	//	nums1[m+i] = nums2[i]
	//}

	//for i := 0; i < m+n-1; i++ {
	//	for j := i + 1; j < m+n; j++ {
	//		if nums1[i] > nums1[j] {
	//			nums1[i], nums1[j] = nums1[j], nums1[i]
	//		}
	//	}
	//}
}

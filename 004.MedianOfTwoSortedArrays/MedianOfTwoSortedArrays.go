package leetcode

// MedianOfTwoSortedArrays 寻找两个有序数组的中位数
func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	mid := totalLength / 2
	// 若数组总长度为奇数，mid==3，则表示取第4个数，所以k=mid+1
	k := mid + 1
	// 若数组总长度为奇数，则取中间那个数
	if totalLength%2 == 1 {
		return float64(getKthElement(nums1, nums2, k))
	}
	// 若数组总长度不为奇数，则取两个中间数的平均数
	return float64(getKthElement(nums1, nums2, k-1)+getKthElement(nums1, nums2, k)) / 2.0
}

// 从两个有序数组获取第 k 个元素
func getKthElement(nums1, nums2 []int, k int) int {
	// p1,p2分别表示当前指针所处的位置
	p1, p2 := 0, 0
	for {
		// 表示nums1已遍历完，从nums2中取即可
		if p1 == len(nums1) {
			return nums2[p2+k-1]
		}
		// 表示nums2已遍历完，从nums1中取即可
		if p2 == len(nums2) {
			return nums1[p1+k-1]
		}
		// 取最小的数时，比较两个数组的第一个数即可
		if k == 1 {
			return min(nums1[p1], nums2[p2])
		}

		// 二分查找，获取数组的第k/2的前一个数，表示可以排除的。需要和数组长度比较，取他们的最小值
		// 因为如果数组长度都没有k/2大，则取数组的最后一个值即可
		// 注意：这里需要加上指针所处的位置值，因为数组的前面几个元素被排除后，指针位置也会变化
		newP1 := min(k/2+p1, len(nums1)) - 1
		newP2 := min(k/2+p2, len(nums2)) - 1
		// 排除一半的数之后，k需要重新赋值，比如原来k==7，排除掉k/2==3个数后，剩下k=7-3=4个数
		// 同时更新指针的位置
		if nums1[newP1] <= nums2[newP2] {
			k = k - (newP1 + 1 - p1)
			p1 = newP1 + 1
		} else {
			k = k - (newP2 + 1 - p2)
			p2 = newP2 + 1
		}
	}

	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//解法一：先合并，再取值，但时间复杂度不满足要求
//func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
//    len1, len2 := len(nums1), len(nums2)
//    p1, p2 := 0, 0
//    var sorted = make([]int, 0, len1+len2)
//    for {
//        if p1 == len1 {
//            sorted = append(sorted, nums2[p2:]...)
//            break
//        }
//        if p2 == len2 {
//            sorted = append(sorted, nums1[p1:]...)
//            break
//        }
//        if nums1[p1] < nums2[p2] {
//            sorted = append(sorted, nums1[p1])
//            p1++
//        } else {
//            sorted = append(sorted, nums2[p2])
//            p2++
//        }
//    }
//
//    mid := (len1+len2)/2
//    if (len1+len2)%2 == 1 {
//        return float64(sorted[mid])
//    }
//
//    return float64(sorted[mid-1]+sorted[mid])/2.0
//}

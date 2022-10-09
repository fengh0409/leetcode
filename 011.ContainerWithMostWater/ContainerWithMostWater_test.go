package leetcode

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	maxArea := ContainerWithMostWater(height)
	if maxArea != expected {
		t.Fatalf("should be %d, but is %d", expected, maxArea)
	}

	// 测试溢出
	height = []int{1, 1}
	expected = 1
	maxArea = ContainerWithMostWater(height)
	if maxArea != expected {
		t.Fatalf("should be %d, but is %d", expected, maxArea)
	}
}

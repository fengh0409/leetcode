package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	result := MaxProfit(prices)
	if result != expected {
		t.Fatalf("should be %d, but is %d", expected, result)
	}

	prices = []int{7, 6, 4, 3, 1}
	expected = 0
	result = MaxProfit(prices)
	if result != expected {
		t.Fatalf("should be %d, but is %d", expected, result)
	}

	prices = []int{2, 5, 1, 3}
	expected = 3
	result = MaxProfit(prices)
	if result != expected {
		t.Fatalf("should be %d, but is %d", expected, result)
	}
}

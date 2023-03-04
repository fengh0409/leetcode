package leetcode

import (
	"fmt"
	"testing"
)

type mock struct {
	input  []int
	output []int
}

var mocks []mock

func setup() {
	mocks = []mock{
		{
			input:  []int{3, 2, 12, 9, 4, 31, 18, 6},
			output: []int{2, 3, 4, 6, 9, 12, 18, 31},
		},
		{
			input:  []int{0, 0, 0, 0, 0, 0},
			output: []int{0, 0, 0, 0, 0, 0},
		},
		{
			input:  []int{0, 0, 1, 0, 0, 0},
			output: []int{0, 0, 0, 0, 0, 1},
		},
		{
			input:  []int{10, 9, 8, 7, 6, 5},
			output: []int{5, 6, 7, 8, 9, 10},
		},
	}
}

func TestBubbleSort(t *testing.T) {
	setup()
	fmt.Printf("BubbleSort unit test:\n")
	for _, item := range mocks {
		var input = make([]int, len(item.input))
		copy(input, item.input)
		result := BubbleSort(item.input)
		fmt.Printf("【input】 : %v\n【output】: %v\n\n", input, result)

	}
}

func TestSelectionSort(t *testing.T) {
	setup()
	fmt.Printf("SelectionSort unit test:\n")
	for _, item := range mocks {
		var input = make([]int, len(item.input))
		copy(input, item.input)
		result := SelectionSort(item.input)
		fmt.Printf("【input】 : %v\n【output】: %v\n\n", input, result)
	}
}

func TestInsertSort(t *testing.T) {
	setup()
	fmt.Printf("InsertSort unit test:\n")
	for _, item := range mocks {
		var input = make([]int, len(item.input))
		copy(input, item.input)
		result := InsertSort(item.input)
		fmt.Printf("【input】 : %v\n【output】: %v\n\n", input, result)
	}
}

func TestMergeSort(t *testing.T) {
	setup()
	fmt.Printf("MergeSort unit test:\n")
	for _, item := range mocks {
		var input = make([]int, len(item.input))
		copy(input, item.input)
		result := MergeSort(item.input)
		fmt.Printf("【input】 : %v\n【output】: %v\n\n", input, result)
	}
}

func TestQuickSort(t *testing.T) {
	setup()
	fmt.Printf("QuickSort unit test:\n")
	for _, item := range mocks {
		var input = make([]int, len(item.input))
		copy(input, item.input)
		result := QuickSort(item.input)
		fmt.Printf("【input】 : %v\n【output】: %v\n\n", input, result)
	}
}

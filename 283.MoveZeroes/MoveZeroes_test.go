package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "Single zero",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "No zeros",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "Zeros at end",
			nums: []int{1, 2, 0, 0},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "Zeros at start",
			nums: []int{0, 0, 1, 2},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "Mixed zeros and negative numbers",
			nums: []int{0, -1, 0, 5, 0, -3},
			want: []int{-1, 5, -3, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.nums))
			copy(got, tt.nums) // Create a copy to avoid modifying the test input directly
			MoveZeroes(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

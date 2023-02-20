package leetcode

import "testing"

func TestNumberOfIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	result := NumIslands(grid)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expected = 3
	result = NumIslands(grid)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}
}

package leetcode

import "testing"

func TestFibonacci(t *testing.T) {
	n := 2
	fn := Fibonacci(n)
	if fn != 1 {
		t.Fatalf("should be 1, but is %d", fn)
	}

	n = 3
	fn = Fibonacci(n)
	if fn != 2 {
		t.Fatalf("should be 2, but is %d", fn)
	}

	n = 4
	fn = Fibonacci(n)
	if fn != 3 {
		t.Fatalf("should be 3, but is %d", fn)
	}

	n = 5
	fn = Fibonacci(n)
	if fn != 5 {
		t.Fatalf("should be 5, but is %d", fn)
	}
}

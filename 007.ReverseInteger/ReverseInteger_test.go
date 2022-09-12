package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	x := 23423523
	expected := 32532432
	rev := ReverseInteger(x)
	if rev != expected {
		t.Fatalf("%d should be %d, but is", expected, rev)
	}

	// 测试溢出
	x = 2147483647
	expected = 0
	rev = ReverseInteger(x)
	if rev != expected {
		t.Fatalf("%d should be %d, but is", expected, rev)
	}

	// 测试溢出
	x = -2147483648
	expected = 0
	rev = ReverseInteger(x)
	if rev != expected {
		t.Fatalf("%d should be %d, but is", expected, rev)
	}
}

package easy

import "testing"

func TestDiagonalDifference(t *testing.T) {
	m := [][]int32{
		{6, 8},
		{-6, 9},
	}

	if r := DiagonalDifference(m); r != 13 {
		t.Error("Should return 13 and is ", r)
	}
}

package easy

import "testing"

func TestFairRotations(t *testing.T) {
	r := fairRations([]int32{1, 2, 3, 4, 5, 6})

	if r != "this" {
		t.Error("should be this and is ", r)
	}
}

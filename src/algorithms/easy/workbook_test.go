package easy

import "testing"

func TestWorkbook(t *testing.T) {
	arr := []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}
	r := workbook(10, 5, arr)

	if r != 4 {
		t.Error("should return 4 and is ", r)
	}
}

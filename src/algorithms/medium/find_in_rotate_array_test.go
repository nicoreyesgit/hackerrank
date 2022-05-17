package medium

import "testing"

func TestFindInRotateArray(t *testing.T) {
	t.Run("should return then", func(t *testing.T) {
		arr := []int{11, 15, 1, 2, 3, 4, 6, 8, 9}
		r := FindInRotateArray(arr, 0, 8, 15)

		if r != 1 {
			t.Error("the index should be 2 and is ", r)
		}

		arr2 := []int{6, 8, 9, 10, 11, 12, 13, 1}
		r = FindInRotateArray(arr2, 0, 7, 1)
		if r != 7 {
			t.Error("the index should be 7 and is ", r)
		}

		arr3 := []int{6, 8, 9, 10, 1, 2, 3, 1}
		r = FindInRotateArray(arr3, 0, 7, 2)
		if r != 5 {
			t.Error("the index should be 7 and is ", r)
		}

		arr4 := []int{1, 2, 3, 5, 7, 8, 10, 11}
		r = FindInRotateArray(arr4, 0, 7, 7)
		if r != 4 {
			t.Error("the index should be 7 and is ", r)
		}

	})

	t.Run("should return -1", func(t *testing.T) {
		arr := []int{6, 8, 9, 10, 11, 12, 13, 1}
		r := FindInRotateArray(arr, 0, 7, 14)
		if r != -1 {
			t.Error("the number should not be in the array")
		}
	})
}

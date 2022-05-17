package easy

import "testing"

func TestFindMissingNumber(t *testing.T) {
	t.Run("should return the number missing", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 8}
		r := FindMissingNumberWithSum(arr)

		if r != 7 {
			t.Error("should return 3 and the return is ", r)
		}
	})
	t.Run("should return the number missing myself", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 8}
		r := FindMissingNumberMyself(arr)

		if r != 7 {
			t.Error("should return 3 and the return is ", r)
		}
	})

}

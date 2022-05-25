package easy

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	ln := len(arr) - 1

	if ln+1 < 2 {
		return 0
	}

	var l, r = arr[0][0], arr[0][ln]

	for i := 1; i <= ln; i++ {
		l += arr[i][i]
		r += arr[i][ln-i]
	}

	return int32(math.Abs(float64(l - r)))
}

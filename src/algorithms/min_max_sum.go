package algorithms

import (
	"fmt"
	"sort"
)

func MinMaxSum(arr []int32) string {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	l := len(arr)
	min := int64(arr[0]) + int64(arr[1]) + int64(arr[2]) + int64(arr[3])
	max := int64(arr[l-1]) + int64(arr[l-2]) + int64(arr[l-3]) + int64(arr[l-4])
	return fmt.Sprintf("%d %d", min, max)
}

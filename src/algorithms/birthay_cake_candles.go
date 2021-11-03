package algorithms

import (
	"sort"
)

func BirthDayCakeCandles(candles []int32) int32 {
	sort.Slice(candles, func(i, j int) bool { return candles[i] < candles[j] })
	count := int32(1)
	num := candles[0]
	var higher int32
	result := make(map[int32]int32, 1)
	for i := 1; i < len(candles); i++ {
		val := candles[i]
		if val == num {
			count++
			if result[val] < count {
				result[val] = count
				higher = val
			}
		} else {
			num = candles[i]
			count = 1
		}
	}
	return result[higher]
}

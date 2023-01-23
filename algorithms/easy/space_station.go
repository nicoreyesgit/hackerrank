package easy

import (
	"math"
	"sort"
)

// https://www.hackerrank.com/challenges/flatland-space-stations/problem?h_r=next-challenge&h_v=zen&isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen
func flatlandSpaceStations(n int32, c []int32) int32 {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	var maxDistance, station, city, si, bDistance, fDistance, bestDistance int32
	lastStation := c[0]
	station = lastStation

	if len(c) == 1 {
		if n > 1 {
			if c[0] > 0 {
				return c[0]
			}
			return n - c[0] - 1
		} else {
			return 0
		}
	}

	for city < n {
		if station == city {
			if si < int32(len(c)) {
				lastStation = c[si]

				si++
				if si <= int32(len(c))-1 && c[si] < n {
					station = c[si]
				}
			}

			city++

			continue
		}

		bDistance = int32(math.Abs(float64(city - lastStation)))
		fDistance = int32(math.Abs(float64(station - city)))
		if bDistance < fDistance {
			bestDistance = bDistance
		} else {
			bestDistance = fDistance
		}

		if bestDistance > maxDistance {
			maxDistance = bestDistance
		}

		city++
	}

	return maxDistance
}

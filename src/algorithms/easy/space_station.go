package easy

import "sort"

// https://www.hackerrank.com/challenges/flatland-space-stations/problem?h_r=next-challenge&h_v=zen&isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen
func flatlandSpaceStations(n int32, c []int32) int32 {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	var maxDistance, topDistance, distance, city, i, j int32

	for city < n-1 {
		i = city

		for j < int32(len(c)) {
			if i == c[j] {
				if city == c[j] {
					j++
				}

				if distance > maxDistance {
					maxDistance = distance
				}

				distance = 0
				city++

				break
			}

			distance++
			i++
		}

		if j == int32(len(c)) && city > c[j-1] {
			topDistance = n - c[j-1] + 1
			break
		}
	}

	if topDistance > maxDistance {
		return topDistance
	}

	return maxDistance
}

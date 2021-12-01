package hard

import (
	"fmt"
	"sort"
)

type replace struct {
	first  []uint64
	second []uint64
}

// KingRichardKnights https://www.hackerrank.com/challenges/king-richards-knights/problem?h_r=next-challenge&h_v=zen
// the first input the large of the array
// the second was modified by me on the template of the HackerRank page and changed to include the commands
// they are the move that the king want to do:
// this size should be 3 and the first value is the line where he want to start to move
// the second value is the first soldier to move
// the third value is how many soldiers (starting from the second value) will move
// knights that the king want to find after the moves
func KingRichardKnights(n uint64, commands [][]uint64, knights []uint64) [][]uint64 {
	// create the knight formation
	initialFormation := knightsArmy(n)
	sort.Slice(knights, func(i, j int) bool {
		return knights[i] < knights[j]
	})
	// this for go through the all commands
	for _, command := range commands {
		line := command[0] - 1
		initialLine := command[0] - 1
		initialIndex := command[1] - 1
		maxIndex := command[2] + command[1] - 1
		// this is to include all cluster of knights that should be moved
		for s := int64(maxIndex - initialIndex); s >= 1; s = int64(maxIndex - initialIndex) {
			rep := replace{
				first:  make([]uint64, s+1),
				second: make([]uint64, s+1),
			}
			r := uint64(0)
			// this is to do the 4 moves of a square
			for i := 0; i < 4; i++ {
				maxValue := maxIndex
				// this is to move the soldiers
				// all the cases should start from the first value to the penultimate
				for j := initialIndex; j < maxIndex; j++ {
					switch i {
					case 0:
						rep.first[r] = initialFormation[line][maxIndex]
						initialFormation[line][maxIndex] = initialFormation[initialLine][j]
					case 1:
						rep.second[r] = initialFormation[maxIndex-1][maxValue]
						initialFormation[maxIndex-1][maxValue] = rep.first[r]
						maxValue--
					case 2:
						rep.first[r] = initialFormation[maxValue-1][initialIndex]
						initialFormation[maxValue-1][initialIndex] = rep.second[r]
						maxValue--
					case 3:
						initialFormation[initialLine][j] = rep.first[r]
						maxValue--
					}
					r++
					line++
				}
				r = 0
				line = initialLine
			}
			initialIndex++
			initialLine++
			line++
			maxIndex--
		}
	}
	return searchValue(initialFormation, knights)
}

func searchValue(army [][]uint64, knights []uint64) [][]uint64 {
	result := make([][]uint64, len(knights))
	for l, line := range army {
		ki, li := 0, 0
		for li < len(line) {
			if ki > len(knights)-1 {
				li++
				ki = 0
			} else if line[li] == knights[ki] {
				result[ki] = []uint64{uint64(l + 1), uint64(li) + 1}
				ki++
				continue
			} else {
				ki++
			}
		}
	}
	return result
}
func knightsArmy(n uint64) [][]uint64 {
	soldiers := make([][]uint64, n)
	var lastV uint64
	for i := uint64(0); i < n; i++ {
		line := make([]uint64, n)
		//go func(index int64) {
		for j := uint64(0); j < n; j++ {
			line[j] = lastV
			lastV++
		}
		soldiers[i] = line
		//}(i)
	}
	return soldiers
}

func cloneArmy(a [][]uint64) [][]uint64 {
	l := len(a)
	r := make([][]uint64, l)
	for i := 0; i < l; i++ {
		newA := make([]uint64, len(a[i]))
		copy(newA, a[i])
		r[i] = newA
	}
	return r
}

func printFormation(l [][]uint64) {
	for j := 0; j < len(l); j++ {
		fmt.Println(l[j])
	}
}

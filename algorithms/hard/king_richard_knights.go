package hard

import (
	"fmt"
)

type replace struct {
	first  []uint64
	second []uint64
}

// KingRichardKnights https://www.hackerrank.com/challenges/king-richards-knights/problem?h_r=next-challenge&h_v=zen
// * the first input the large of the array
// * the second was modified by me on the template of the HackerRank page and changed to include the commands
// * they are the move that the king want to do:
//   - this length should be 3 and the first value is the line where he want to start to move
//   - the second value is the first soldier to move
//   - the third value is how many soldiers (starting from the second value) will be moved
//
// * knights = kinghts that the king want to find after the moves
func KingRichardKnights(n uint64, commands [][]uint64, knights []uint64) [][]uint64 {
	// create the knight formation
	initialFormation := knightsArmy(n)
	// this for go through the all commands
	for _, command := range commands {
		line := command[0] - 1
		initialLine := command[0] - 1
		initialIndex := command[1] - 1
		maxIndex := command[2] + command[1] - 1
		// this is to include all cluster of knights that should be moved
		for t := uint64(0); t < maxIndex/2; t++ {
			s := int64(maxIndex - initialIndex)
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
	count := 0
	for l, line := range army {
		ki, li := 0, 0
		for li < len(line) {
			count++
			if ki > len(knights)-1 {
				li++
				ki = 0
				// TODO: improve the algorithm to exclude in the search the knights already found
			} else if line[li] == knights[ki] {
				result[ki] = []uint64{uint64(l + 1), uint64(li) + 1}
				ki++
				continue
			} else {
				ki++
			}
		}
	}
	fmt.Println(count)
	return result
}
func knightsArmy(n uint64) [][]uint64 {
	soldiers := make([][]uint64, n)
	var lastV uint64
	// TODO: improve the algorithm to remove one 'for'
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

func printFormation(l [][]uint64) {
	for j := 0; j < len(l); j++ {
		fmt.Println(l[j])
	}
}

package hard

import "fmt"

// KingRichardKnights https://www.hackerrank.com/challenges/king-richards-knights/problem?h_r=next-challenge&h_v=zen
// the first input the large of the array
// the second was modified by me on the template of the HackerRank page and changed to include the commands
// they are the move that the king want to do:
// this size should be 3 and the first value is the line where he want to start to move
// the second value is the first soldier to move
// the third value is how many soldiers (starting from the second value) will move
// knights that the king want to find after the moves
func KingRichardKnights(n int32, commands [][]int32, knights []int32) [][]int32 {
	// create the knight formation
	initialFormation := knightsArmy(n)
	fmt.Println("the initial formation")
	printFormation(initialFormation)
	// this for go through the all commands
	for move := 0; move < len(commands); move++ {
		// clone the army
		newFormation := cloneArmy(initialFormation)
		command := commands[move]
		line := command[0] - 1
		initialLine := command[0] - 1
		initialIndex := command[1] - 1
		maxIndex := command[2] + command[1] - 1
		// this is to include all sub group of knights that should be moved
		for s := maxIndex - initialIndex; s >= 1; s = maxIndex - initialIndex {
			// this is to do the 4 moves of a square
			for i := 0; i < 4; i++ {
				maxValue := maxIndex
				// this is to move the soldiers
				// all the cases should start from the first value to the penultimate
				for j := initialIndex; j < maxIndex; j++ {
					switch i {
					case 0:
						newFormation[line][maxIndex] = initialFormation[initialLine][j]
					case 1:
						newFormation[maxIndex-1][maxValue] = initialFormation[line][maxIndex]
						maxValue--
					case 2:
						newFormation[maxValue-1][initialIndex] = initialFormation[maxIndex-1][maxValue]
						maxValue--
					case 3:
						newFormation[initialLine][j] = initialFormation[maxValue-1][initialIndex]
						maxValue--
					}
					line++
				}
				line = initialLine
			}
			initialIndex++
			initialLine++
			line++
			maxIndex--
		}
		initialFormation = newFormation
	}
	fmt.Println("the new formation is")
	printFormation(initialFormation)
	return initialFormation
}

func knightsArmy(n int32) [][]int32 {
	soldiers := make([][]int32, n)
	var lastV int32
	for i := int32(0); i < n; i++ {
		line := make([]int32, n)
		for j := int32(0); j < n; j++ {
			line[j] = lastV
			lastV++
		}
		soldiers[i] = line
	}
	return soldiers
}

func printFormation(l [][]int32) {
	for j := 0; j < len(l); j++ {
		fmt.Println(l[j])
	}
}

func cloneArmy(a [][]int32) [][]int32 {
	l := len(a)
	r := make([][]int32, l)
	for i := 0; i < l; i++ {
		newA := make([]int32, len(a[i]))
		copy(newA, a[i])
		r[i] = newA
	}
	return r
}

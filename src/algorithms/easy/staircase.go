package easy

import "fmt"

func StairCase(stairNumber int32) []string {
	stairParsed := int(stairNumber)
	var result []string
	if stairNumber < 0 || stairNumber > 100 {
		return result
	}
	for i := 1; i <= stairParsed; i++ {
		spaces := padding(stairParsed - i)
		hashes := hashes(i - 1)
		line := spaces + hashes
		fmt.Println(line)
		result = append(result, line)
	}
	return result
}

func hashes(qty int) string {
	hashes := ""
	for i := 0; i <= qty; i++ {
		hashes += "#"
	}
	return hashes
}

func padding(spaces int) string {
	padded := ""
	for i := 0; i < spaces; i++ {
		padded += " "
	}
	return padded
}

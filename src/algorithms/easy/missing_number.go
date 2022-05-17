package easy

func FindMissingNumberWithSum(arr []int) int {
	var (
		sumArr int
		l      = len(arr) + 1
	)
	sum := (l * (l + 1)) / 2

	for i := 0; i < len(arr); i++ {
		sumArr += arr[i]
	}

	return sum - sumArr
}

func FindMissingNumberMyself(arr []int) int {
	aux := 1

	for i := 0; i < len(arr); i++ {
		if arr[i] != aux {
			return aux
		}
		aux++
	}

	return 0
}

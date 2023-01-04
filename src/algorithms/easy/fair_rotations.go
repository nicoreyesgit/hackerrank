package easy

func fairRations(b []int32) string {
	// Write your code here
	ln := len(b)
	// l := b[0]

	if ln == 1 {
		if b[0]%2 != 0 || b[1]%2 != 0 {
			return "NO"
		}
	}

	return ""
}

package medium

func FindInRotateArray(arr []int, l, r, v int) int {
	m := l + (r-l)/2

	if m <= len(arr)-1 {
		if arr[m] == v {
			return m
		}

		if arr[l] > arr[m] && (arr[l] <= v && arr[m] < v) {
			return FindInRotateArray(arr, l, m-1, v)
		}

		if arr[r] < arr[m] && (arr[m] > v && arr[r] <= v) {
			return FindInRotateArray(arr, m+1, r, v)
		}

		if arr[m] > v {
			return FindInRotateArray(arr, l, m-1, v)
		}

		return FindInRotateArray(arr, m+1, r, v)
	}

	return -1
}

package medium

func FindInRotateArrayRecursive(arr []int, l, r, v int) int {
	m := l + (r-l)/2

	if m <= len(arr)-1 {
		if arr[m] == v {
			return m
		}

		if arr[l] > arr[m] && (arr[l] <= v && arr[m] < v) {
			return FindInRotateArrayRecursive(arr, l, m-1, v)
		}

		if arr[r] < arr[m] && (arr[m] > v && arr[r] <= v) {
			return FindInRotateArrayRecursive(arr, m+1, r, v)
		}

		if arr[m] > v {
			return FindInRotateArrayRecursive(arr, l, m-1, v)
		}

		return FindInRotateArrayRecursive(arr, m+1, r, v)
	}

	return -1
}

func FindInRotateArrayLoop(arr []int, l, r, v int) int {

	for l < r {
		m := l + (r-l)/2
		if arr[m] == v {
			return m
		}

		if arr[l] > arr[m] && (arr[l] <= v && arr[m] < v) {
			r = m - 1
		}

		if arr[r] < arr[m] && (arr[m] > v && arr[r] <= v) {
			l = m + 1
		}

		if arr[m] > v {
			r = m - 1
		}

		l = m + 1
	}

	return -1
}

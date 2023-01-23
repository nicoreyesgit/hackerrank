package easy

// https://www.hackerrank.com/challenges/lisa-workbook/problem?h_r=next-challenge&h_v=zen&isFullScreen=true&h_r=next-challenge&h_v=zen
// n number of chapters on workbook
// k max number of problems that a page can hold
// array indicate how many problems has a chapter
// ex: arr =  [1,2,3,4,5], n = 5 so
// chapter 1 = 1 problem
// chapter 2 = 2 problems
// etc
func workbook(n int32, k int32, arr []int32) int32 {
	// Write your code here
	var (
		specials      int32
		problemNumber = int32(1)
		page          = int32(1)
		remanent      = arr[0]
		auxK          = k
		i             = int32(0)
	)

	for i < n {
		if problemNumber == page {
			specials++
		}

		if problemNumber == arr[i] {
			problemNumber = 1
			page++
			i++

			if i < n {
				remanent = arr[i]
				if remanent >= k {
					auxK = k
				} else {
					auxK = remanent
				}
			}

			continue
		}

		if auxK == 1 {
			remanent = remanent - k

			if remanent >= k {
				auxK = k
				page++
				problemNumber++
			} else if remanent > 0 {
				auxK = remanent
				page++
				problemNumber++
			} else {
				problemNumber = 1
				i++
				page++

				if i < n {
					remanent = arr[i]
					if remanent >= k {
						auxK = k
					} else {
						auxK = remanent
					}
				}
			}

			continue
		}

		auxK--
		problemNumber++
	}

	return specials
}

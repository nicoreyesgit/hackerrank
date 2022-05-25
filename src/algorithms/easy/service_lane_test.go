package easy

import (
	"fmt"
	"testing"
)

func TestServiceLine(t *testing.T) {
	cases := [][]int32{{2, 3}, {1, 4}, {2, 4}, {2, 4}, {2, 3}}
	width := []int32{1, 2, 2, 2, 1}

	ex := []int32{2, 1, 1, 1, 2}
	r := serviceLane(8, width, cases)

	for i := 0; i < len(r); i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if r[i] != ex[i] {
				t.Errorf("%d should be 1 and is %v", i, r[i])
			}

		})
	}
}

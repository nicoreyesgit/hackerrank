package hard

import (
	"testing"
)

func TestKinRichardKnights(t *testing.T) {
	t.Run("should return a two dimenssion array", func(t *testing.T) {
		result := KingRichardKnights(int32(7), [][]int32{
			{1, 2, 4},
			{2, 3, 3},
			{3, 4, 1},
			{3, 4, 0},
		}, []int32{1, 2, 4})
		if result[0][1] != int32(1) {
			t.Errorf("should return 1 and is %v", result[0][1])
		}
		if result[2][2] != int32(10) {
			t.Errorf("should return 1 and is %v", result[2][2])
		}
	})
}

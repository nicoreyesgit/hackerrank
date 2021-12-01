package hard

import (
	"testing"
)

func TestKinRichardKnights(t *testing.T) {
	t.Run("should return a two dimension array", func(t *testing.T) {
		result := KingRichardKnights(uint64(7), [][]uint64{
			{1, 2, 4},
			{2, 3, 3},
			{3, 4, 1},
			{3, 4, 0},
		}, []uint64{0,
			9,
			6,
			11,
			48,
			25,
			24})
		if result[0][1] != uint64(1) {
			t.Errorf("should return 1 and is %v", result[0][1])
		}
		if result[5][1] != uint64(4) {
			t.Errorf("should return 1 and is %v", result[0][1])
		}
	})
}

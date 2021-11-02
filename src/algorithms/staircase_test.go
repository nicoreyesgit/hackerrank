package algorithms

import "testing"

func TestStairCase(t *testing.T) {
	t.Run("when the input is valid should return the stair", func(t *testing.T) {
		result := StairCase(5)
		if result[0] != "    #" {
			t.Errorf("Should return empty but return %v", result)
		}
		if result[4] != "#####" {
			t.Errorf("Should return empty but return %v", result)
		}
	})
}

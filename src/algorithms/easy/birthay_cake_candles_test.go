package easy

import "testing"

func TestBirthDayCakeCandles(t *testing.T) {
	t.Run("when the element 3 is high and it repeat 3 times should return 3", func(t *testing.T) {
		candles := []int32{1, 2, 3, 4, 3, 3}
		result := BirthDayCakeCandles(candles)
		if result != 3 {
			t.Errorf("Should return 3 and the result is %d", result)
		}
	})
}

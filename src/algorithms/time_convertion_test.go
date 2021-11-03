package algorithms

import "testing"

func TestTimeConversion(t *testing.T) {
	t.Run("when the input is 12 AM time should return 00", func(t *testing.T) {
		result := TimeConversion("12:45:54AM")
		t.Logf("the result is %v", result)
		if result != "00:45:54" {
			t.Errorf("Should return 00:45:54 and return %v", result)
		}
	})
}

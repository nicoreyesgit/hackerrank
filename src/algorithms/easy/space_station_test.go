package easy

import "testing"

func TestFlatlandSpaceStations(t *testing.T) {
	stations := []int32{0, 4}
	r := flatlandSpaceStations(5, stations)

	if r != 2 {
		t.Error("should be 0 and is ", r)
	}
}

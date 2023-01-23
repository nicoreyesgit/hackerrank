package easy

import "testing"

func TestFlatlandSpaceStations(t *testing.T) {
	stations := []int32{99}
	r := flatlandSpaceStations(100, stations)

	if r != 20 {
		t.Error("should be 0 and is ", r)
	}
}

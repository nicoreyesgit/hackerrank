package medium

import (
	"testing"
)

func TestFindPathWithRecursion(t *testing.T) {
	way := []string{
		"#####################",
		"#            ########",
		"# ##########E########",
		"# ###################",
		"# ###################",
		"#         ###########",
		"#########S###########",
	}
	endPoint := FindPathWithRecursion(way, Point{9, 6}, Point{12, 2})
	last := endPoint[len(endPoint)-1]

	if last.X == 12 && last.Y == 2 {
		t.Log("There is no error")
	} else {
		t.Errorf("should find the path: %v", last)
	}
}

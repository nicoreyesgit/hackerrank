package algorithms

import "testing"

func TestMinMaxSum(t *testing.T) {
	t.Run("when the input is 1,2,3,4,5 should return the min sum of four elements and the max", func(t *testing.T) {
		arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
		result := MinMaxSum(arr)
		if result != "2063136757 2744467344" {
			t.Errorf("The result should be 2063136757  2744467344 and is %v", result)
		}
		arr2 := []int32{3, 5, 7, 13, 18}
		result2 := MinMaxSum(arr2)
		if result2 != "28 43" {
			t.Errorf("The result should be  and is %v", result2)
		}
	})
}

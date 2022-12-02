package aoc

import (
	"fmt"
	"testing"
)

func TestMinMaxTestMinMax(t *testing.T) {
	testCases := []struct {
		arr []int
		min int
		max int
	}{
		{[]int{5}, 5, 5},
		{[]int{9, 1}, 1, 9},
		{[]int{1, 9}, 1, 9},
		{[]int{32, 23, 999, -45}, -45, 999},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v, %v\n", tc.arr, tc.min, tc.max), func(t *testing.T) {
			min, max := MinMax(tc.arr)
			if min != tc.min {
				t.Errorf("got min %v, want min %v", min, tc.min)
			}
			if max != tc.max {
				t.Errorf("got max %v, want max %v", max, tc.max)
			}
		})
	}
}

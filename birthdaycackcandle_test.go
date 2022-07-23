package main

import (
	"fmt"
	"testing"
)

func TestCountTallestCandles(t *testing.T) {
	//arrange
	testcases := []struct {
		in  []int32
		out int32
	}{
		{[]int32{4, 4, 1, 3}, 2},
		{[]int32{3, 2, 1, 3}, 2},
		{[]int32{6, 7, 6, 5, 9}, 1},
		{[]int32{3, 6, 1, 8, 7}, 1},
		{[]int32{10, 3, 8, 1, 9, 1}, 1},
		{[]int32{7, 5, 6, 5, 8, 8}, 2},
		{[]int32{7, 5, 9, 5, 9, 9}, 3},
	}

	for _, test := range testcases {
		t.Run("test", func(t *testing.T) {
			//act
			res := CountCandles(test.in)

			//assert
			if test.out != res {
				t.Error(fmt.Sprintf("the count tallest of candles %v should be %v but have %v", test.in, test.out, res))
			}
		})
	}
}

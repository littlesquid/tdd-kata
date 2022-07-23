package main

import (
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	testcases := []struct {
		in  []int32
		out string
	}{
		{[]int32{1, 2, 3, 4, 5}, "10 14"},
		{[]int32{1, 3, 5, 7, 9}, "16 24"},
		{[]int32{396285104, 573261094, 759641832, 819230764, 364801279}, "2093989309 2548418794"},
		{[]int32{140638725, 436257910, 953274816, 734065819, 362748590}, "1673711044 2486347135"},
	}

	for _, test := range testcases {
		t.Run(test.out, func(t *testing.T) {
			res := MiniMaxSum(test.in)
			if test.out != res {
				Assert(t, test.out, res)
			}
		})
	}
}

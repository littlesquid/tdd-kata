package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStudentGrading(t *testing.T) {
	//arrange
	testcases := []struct {
		in  []int32
		out []int32
	}{
		{[]int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
		{[]int32{84, 29, 57}, []int32{85, 29, 57}},
	}

	for _, test := range testcases {
		t.Run("test", func(t *testing.T) {
			//act
			res := StudentGrading(test.in)

			//assert
			if !reflect.DeepEqual(test.out, res) {
				t.Error(fmt.Sprintf("the grade after rouding as appropriate of %v should be %v but have %v", test.in, test.out, res))
			}
		})
	}
}

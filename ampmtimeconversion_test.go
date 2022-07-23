package main

import (
	"fmt"
	"testing"
)

func TestTimeConversionAMPM(t *testing.T) {
	//arrange
	testcases := []struct {
		in  string
		out string
	}{
		{"12:01:00PM", "12:01:00"},
		{"05:32:52AM", "05:32:52"},
		{"12:01:00AM", "00:01:00"},
		{"07:05:45PM", "19:05:45"},
		{"01:23:22PM", "13:23:22"},
		{"05:05:45", "invalid input format"},
		{"17:35:45PM", "17:35:45"},
		{"17:35:45AM", "invalid input format"},
	}

	for _, test := range testcases {
		t.Run("timeconversiontest", func(t *testing.T) {
			//act
			res := TimeConversion(test.in)

			//assert
			if test.out != res {
				t.Error(fmt.Sprintf("time conversion of %v should be %v but have %v", test.in, test.out, res))
			}
		})
	}

}

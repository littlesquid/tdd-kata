package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func TimeConversion(time string) string {
	matchPM, _ := regexp.MatchString(`((\d{2})|0[0-9]):((\d{2})|0[0-9]):((\d{2})|0[0-9])(PM|pm)$`, time)
	matchAM, _ := regexp.MatchString(`((\d{2})|0[0-9]):((\d{2})|0[0-9]):((\d{2})|0[0-9])(AM|am)$`, time)

	replacer := strings.NewReplacer("AM", "", "PM", "")
	times := strings.Split(time, ":")

	hh, _ := strconv.ParseInt(times[0], 10, 32)
	mm := times[1]
	ss := replacer.Replace(times[2])

	if matchAM {
		if hh > 12 {
			return "invalid input format"
		}
		if hh == 12 {
			hh = 0
		}
		return fmt.Sprintf("%02d:%v:%v", hh, mm, ss)
	}

	if matchPM {
		if hh < 12 {
			hh += 12
		}
		return fmt.Sprintf("%02d:%v:%v", hh, mm, ss)
	}
	return "invalid input format"
}

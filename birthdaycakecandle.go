package main

import "fmt"

func CountCandles(candles []int32) int32 {
	candlesToCount := make(map[int32]int32)
	for _, c := range candles {
		candlesToCount[c] += 1
	}
	fmt.Println(fmt.Sprintf(">>>>>%v", candlesToCount))

	max := int32(0)
	for key := range candlesToCount {
		if key > max {
			max = key
		}
	}
	count := candlesToCount[max]
	fmt.Println(max)
	fmt.Println(max)
	return count

}

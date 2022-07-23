package main

import "fmt"

func MiniMaxSum(arr []int32) string {
	sum := int64(0)
	min := int64(0)
	max := int64(0)

	for round := 0; round < len(arr); round++ {
		for i, number := range arr {
			if i != round {
				sum += int64(number)
			}
			if i == len(arr)-1 && min == 0 {
				min = sum
			}
		}
		if sum < min {
			min = sum
		} else if sum > max {
			max = sum
		}
		sum = 0
	}
	return fmt.Sprintf("%v %v", min, max)
}

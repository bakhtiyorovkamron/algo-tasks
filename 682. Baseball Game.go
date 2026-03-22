package main

import (
	"strconv"
)

func CalPoints(operations []string) int {

	var nums []int
	sum := 0
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
		case "D":
			nums = append(nums, nums[len(nums)-1]*2)
		case "C":
			nums = nums[0 : len(nums)-1]
		default:
			n, _ := strconv.Atoi(operations[i])
			nums = append(nums, n)
		}
	}

	for i := range nums {
		sum += nums[i]
	}

	return sum
}

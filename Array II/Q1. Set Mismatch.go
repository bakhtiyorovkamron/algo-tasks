package main

import "fmt"

// sort it
// remove duplicates
// calculate

func FindErrorNums(nums []int) []int {
	fmt.Println("INPUT :",nums)
	for i := range nums {

		if i != len(nums)-1 && nums[i] == nums[i+1] {
			if nums[0] >= nums[i+1] {
				return []int{nums[i], nums[i+1] - 1}
			} else {
				return []int{nums[i], nums[i+1] + 1}
			}
		}

	}
	return []int{}
}

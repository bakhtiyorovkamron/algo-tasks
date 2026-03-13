package main

import (
	"fmt"
	"sort"
)

func FindDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)
	fmt.Println("Input :", nums)
	outPut := []int{}

	clearedNums := make(map[int]struct{})
	for i := range nums {
		clearedNums[nums[i]] = struct{}{}
	}
	for 

	return outPut
}

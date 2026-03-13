package main

import (
	"fmt"
	"sort"
)

func FindDisappearedNumbers(nums []int) []int {

	fmt.Println("Input :", nums)
	outPut := []int{}

	clearedNums := make(map[int]struct{})
	for i := range nums {
		if _, ok := clearedNums[nums[i]]; ok {
			continue
		}
		clearedNums[nums[i]] = struct{}{}
		outPut = append(outPut, nums[i])
	}

	sort.Ints(outPut)

	for i := range outPut {
		if _, ok := clearedNums[i+1]; !ok {
			fmt.Println(i+1)
		}
	}

	return outPut
}

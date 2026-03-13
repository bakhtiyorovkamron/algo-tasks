package main

import (
	"sort"
)

func FindDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)

	duplc := make(map[int]struct{})
	for i := range nums {
		duplc[nums[i]] = struct{}{}
	}

	result := []int{}

	for i := range nums {
		if _, ok := duplc[i+1]; !ok {
			result = append(result, i+1)
		}

	}
	return result
}

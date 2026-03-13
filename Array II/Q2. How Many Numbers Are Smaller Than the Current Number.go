package main

import (
	"fmt"
	"sort"
)

func SmallerNumbersThanCurrentf(nums []int) []int {
	length := len(nums)
	outPut := []int{}
	for i := 0; i < length; i++ {
		n := nums[i]
		smallestThanNCount := 0
		for j := 0; j < length; j++ {
			if n > nums[j] {
				smallestThanNCount++
			}
		}
		outPut = append(outPut, smallestThanNCount)
	}

	return outPut
}

func SmallerNumbersThanCurrent(nums []int) []int {
	n := make([]int, len(nums), len(nums))
	copy(n, nums)

	outPut := []int{}

	sort.Ints(nums)

	f := make(map[int]int)

	for _, v := range nums {
		f[v]++
	}
	a := 0
	for _, n := range nums {
		a += f[n]
		fmt.Println(n, a)
	}

	return outPut
}

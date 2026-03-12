package main

func SmallerNumbersThanCurrent(nums []int) []int {
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

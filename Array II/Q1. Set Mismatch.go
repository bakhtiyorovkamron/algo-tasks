package main

// sort it
// remove duplicates
// calculate

func FindErrorNums(nums []int) []int {

	duplicates := make(map[int]struct{})
	numsWithoutDupls := []int{}
	l := len(nums)
	d := 0
	for i := range nums {
		if _, ok := duplicates[nums[i]]; ok {
			d = nums[i]
			continue
		}
		duplicates[nums[i]] = struct{}{}
		numsWithoutDupls = append(numsWithoutDupls, nums[i])
	}
	for i := 0; i < l; i++ {
		if _, ok := duplicates[i+1]; !ok {
			return []int{d, i + 1}
		}
	}

	return []int{}
}

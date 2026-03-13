package main

func FindMaxConsecutiveOnes(nums []int) int {
	//1,0,1,1,0,1
	m := 0
	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 1 {
			m = 0
			continue
		}
		m++
		if a < m {
			a = m
		}
	}

	return a
}

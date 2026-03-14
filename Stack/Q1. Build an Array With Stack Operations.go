package main

func BuildArray(target []int, n int) []string {
	outPut := []string{}
	i := 1
	index := 0
	for ; i <= n; i++ {
		outPut = append(outPut, "Push")
		if target[index] != i {
			outPut = append(outPut, "Pop")
		} else {
			if index != len(target)-1 {
				index++
			}
		}

		if target[index] == i {
			return outPut
		}

	}

	return outPut
}

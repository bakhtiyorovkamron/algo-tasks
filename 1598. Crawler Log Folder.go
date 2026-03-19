package main

func MinOperations(logs []string) int {
	a := 0
	for _, v := range logs {
		if v == "../" {
			if a > 0 {
				a--
			}
		} else if v != "./" {
			a++
		}
	}

	if a < 0 {
		return 0
	}

	return a
}

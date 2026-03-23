package main

func MaxDepth(s string) int {
	a := 0
	depth := 0
	for i := range s {
		switch s[i] {
		case '(':
			depth++
		case ')':
			if depth > a {
				a = depth
			}
			depth--
		}
	}
	return a
}

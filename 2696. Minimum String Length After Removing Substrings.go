package main

func MinLength(s string) int {
	ab := "AB"
	cd := "CD"

	stack := []string{}

	for _, v := range s {
		stack = append(stack, string(v))
	}
	i := 0

	for i < len(stack) {
		if i+1 == len(stack) {
			break
		}
		abOrcd := stack[i] + stack[i+1]
		if abOrcd == ab || abOrcd == cd {
			stack = append(stack[0:i], stack[i+2:]...)
			i = 0
			continue
		}
		i++

	}

	return len(stack)
}

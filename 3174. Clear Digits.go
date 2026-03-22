package main

func ClearDigits(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

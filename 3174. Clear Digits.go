package main

func ClearDigits(s string) string {
	a := "0123456789"

	numbers := make(map[uint8]struct{})
	for i := 0; i < len(a); i++ {
		numbers[a[i]] = struct{}{}
	}

	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if _, ok := numbers[b[i]]; ok {
			b = append(b[:i-1], b[i+1:]...)
			i = 0
		}
	}
	s = string(b)
	return s
}

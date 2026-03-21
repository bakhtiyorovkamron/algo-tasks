package main

import (
	"strings"
)

func MakeGood(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {

		if len(stack) != 0 {
			if strings.ToUpper(string(stack[len(stack)-1])) == strings.ToUpper(string(s[i])) && string(stack[len(stack)-1]) != string(s[i]) {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	return string(stack)
}

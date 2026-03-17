package main

import (
	"strconv"
)

type stack []int

func (s *stack) pop() int {
	lastIndex := len(*s) - 1
	val := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return val
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func EvalRPN(tokens []string) int {

	stack := &stack{}

	for _, v := range tokens {

		switch v {
		case "+":
			lastElement := stack.pop()
			beforLastestElement := stack.pop()

			stack.push(beforLastestElement + lastElement)
		case "-":
			lastElement := stack.pop()
			beforLastestElement := stack.pop()

			stack.push(beforLastestElement - lastElement)
		case "*":
			lastElement := stack.pop()
			beforLastestElement := stack.pop()

			stack.push(beforLastestElement * lastElement)
		case "/":
			lastElement := stack.pop()
			beforLastestElement := stack.pop()

			stack.push(beforLastestElement / lastElement)
		default:
			num, _ := strconv.Atoi(v)
			stack.push(num)
		}

	}

	return stack.pop()
}

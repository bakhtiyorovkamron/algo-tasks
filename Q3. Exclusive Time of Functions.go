package main

import (
	"fmt"
	"strings"
)

type Log struct {
	Name      string
	Operation string
	Time      string
}

type Stack map[string]int

func NewStack() Stack {
	return make(map[string]int)
}

func (s Stack) push(x string) {
	if _, ok := s[x]; !ok {
		s = map[string]int{}
	}
	s[x]++
}
func (s Stack) pop(x string) int {
	t := s[x]
	s = map[string]int{}
	return t
}

func ExclusiveTime(n int, logs []string) []int {

	stack := NewStack()
	stack.push("1")
	stack.push("1")
	fmt.Println(stack)

	for _, log := range logs {
		logParser(log)
	}

	return []int{}
}
func logParser(log string) Log {
	l := strings.Split(log, ":")
	if len(l) != 3 {
		return Log{}
	}
	return Log{
		Name:      l[0],
		Operation: l[1],
		Time:      l[2],
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	Name      string
	Operation string
	At        string
}

type Stack []string

func empty() Stack {
	return []string{}
}

var result []string

var stackMap = make(map[string][]string)

func (s *Stack) push(x Log) {

	if x.Operation == "start" {
		at, _ := strconv.Atoi(x.At)
		if _, ok := stackMap[x.Name]; !ok {
			if len(*s) != 0 {
				at--
				stackMap[(*s)[len(*s)-1]] = append(stackMap[(*s)[len(*s)-1]], x.At)
			}
		}
		*s = append(*s, x.Name)
		stackMap[x.Name] = append(stackMap[x.Name], fmt.Sprintf("%d", at))
	} else {
		stackMap[(*s)[len(*s)-1]] = append(stackMap[(*s)[len(*s)-1]], x.At)
		*s = (*s)[:len(*s)-1]

		if len(*s) != 0 {
			stackMap[(*s)[len(*s)-1]] = append(stackMap[(*s)[len(*s)-1]], x.At)
		}
	}
}

func (s *Stack) remove(x string, times []string) {
	//for _, t := range times {
	//
	//}
}

func ExclusiveTime(n int, logs []string) []int {
	fmt.Println("LOGS :", logs)
	//INPUT : "0:start:0","1:start:2","1:end:5","0:end:6"
	//OUTPUT : [3,4]
	stack := Stack{}

	for _, log := range logs {
		l := logParser(log)
		stack.push(l)

	}
	fmt.Println(stackMap)
	fmt.Println(stack)
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
		At:        l[2],
	}
}

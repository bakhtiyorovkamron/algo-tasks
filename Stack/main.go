package main

import "fmt"

func main() {
	fmt.Println(EvalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}

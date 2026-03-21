package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {

	nodeValues := []int{}

	for head.Next != nil {
		nodeValues = append(nodeValues, head.Val)
		head = head.Next
	}
	nodeValues = append(nodeValues, head.Val)

	for i := 0; i < len(nodeValues)/2; i++ {
		if nodeValues[i] != nodeValues[len(nodeValues)-1-i] {
			return false
		}
	}

	return true
}

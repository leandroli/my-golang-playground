package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result []int
	p := head
	for p != nil {
		result = append([]int{p.Val}, result...)
		p = p.Next
	}
	return result
}

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 =  ListNode{0, nil}
	var node1, node2, node3 = l1, l2, &l3
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			node3.Next = &ListNode{node1.Val, nil}
			node1 = node1.Next
		} else {
			node3.Next = &ListNode{node2.Val, nil}
			node2 = node2.Next
		}
		node3 = node3.Next
	}
	if node1 == nil {
		for node2 != nil {
			node3.Next = &ListNode{node2.Val, nil}
			node2 = node2.Next
			node3 = node3.Next
		}
	}

	if node2 == nil {
		for node1 != nil {
			node3.Next = &ListNode{node1.Val, nil}
			node1 = node1.Next
			node3 = node3.Next
		}
	}
	return l3.Next
}


func main() {
	l1 := &ListNode{Val:1}
	l1.Next = &ListNode{Val:3}
	l1.Next.Next =&ListNode{Val:5}
	l2 := &ListNode{Val:1}
	l2.Next = &ListNode{Val:4}
	l2.Next.Next = &ListNode{Val:6}

	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3, l3.Next, l3.Next.Next)
}

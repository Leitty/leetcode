package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	first := &ListNode{0, nil}
	first.Next = head
	var node1 , node2 , node3 *ListNode
	node1 = first
	for node1 != nil && node1.Next != nil && node1.Next.Next != nil {
		node2 = node1.Next
		node3 = node2.Next
		node1.Next, node3.Next, node2.Next = node3, node2, node3.Next
		node1 = node2
	}
	return first.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3, nil}
	l1.Next.Next.Next = &ListNode{4, nil}
	//var l1 *ListNode
	//l1 = nil
	l2 := swapPairs(l1)
	printList(l2)
}

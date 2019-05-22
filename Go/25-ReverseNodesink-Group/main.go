package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var node1, node2, node3 *ListNode
	var n int
	for node1 = head; node1 != nil; {
		node1 = node1.Next
		n++
	}
	if n < k {
		return head
	}

	first := &ListNode{0, head}
	node1 = first
	var term int
	for term = n/k;term != 0;term-- {
		node2, node3 = reverseListNode(node1.Next, k)
		node1.Next = node2
		node1 = node3
	}
	return first.Next
}

func reverseListNode(head *ListNode, k int)  (*ListNode, *ListNode){
	var node1, node2, node3 *ListNode
	node1 =head
	node2 = node1.Next
	for n:=k-1; n!=0; n-- {
		node3 = node2.Next
		node2.Next = node1
		node1, node2 = node2, node3
	}
	head.Next = node2
	return node1, head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	l1 := &ListNode{1,nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3,nil}
	l1.Next.Next.Next = &ListNode{4, nil}
	l1.Next.Next.Next.Next = &ListNode{5, nil}
	printList(reverseKGroup(l1,3))
}

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var NodeNow *ListNode = new(ListNode)
//	var l3 *ListNode
//	var NodeThree *ListNode
//	//NodeTwo  := l2
//	mode := 0
//	for NodeOne,NodeTwo := l1, l2; NodeOne != nil || NodeTwo !=nil; NodeOne = NodeOne.Next {
//		sum := NodeOne.Val + NodeTwo.Val
//		NodeNow.Val = sum % 10 + mode
//		if mode != 0 {
//			mode = 0
//		}
//		mode = sum / 10
//		NodeTwo = NodeTwo.Next
//		if l3 == nil {
//			l3 = new(ListNode)
//			l3.Val = NodeNow.Val
//			NodeThree = l3
//		}else {
//			if NodeThree.Next == nil {
//				NodeThree.Next = new(ListNode)
//			}
//			NodeThree.Next.Val = NodeNow.Val
//			NodeThree = NodeThree.Next
//		}
//	}
//	return l3
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, mode int
	//l3 := new(ListNode)
	l3 := ListNode{0,nil}
	NodeNow := &l3
	for NodeOne, NodeTwo := l1, l2; NodeOne != nil || NodeTwo !=nil || mode !=0;  {
		v1, v2 := 0, 0
		if NodeOne != nil {
			v1 = NodeOne.Val
			NodeOne = NodeOne.Next
		}
		if NodeTwo != nil {
			v2 = NodeTwo.Val
			NodeTwo = NodeTwo.Next
		}

		sum = v1 + v2 + mode
		NodeNow.Next = &ListNode{sum%10,nil}
		NodeNow = NodeNow.Next
		mode = sum/10
	}
	return l3.Next
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := ListNode{0, nil}
//	current := &head
//	p := l1;
//	q := l2;
//	carry := 0
//	for {
//		v1 := 0
//		v2 := 0
//		if p != nil {
//			v1 = p.Val
//			p = p.Next;
//		}
//		if q != nil {
//			v2 = q.Val
//			q = q.Next;
//		}
//		sum := carry + v1 + v2
//		carry = sum / 10
//		current.Next = &ListNode{sum % 10, nil}
//		current = current.Next
//		if p == nil && q == nil {
//			break
//		}
//	}
//	if carry != 0 {
//		current.Next = &ListNode{carry, nil}
//	}
//	return head.Next
//}


func main() {
	//l1 := &ListNode {Val:2}
	//l1.Next = &ListNode{Val:4}
	//l1.Next.Next = &ListNode{Val:3}
	//
	//l2 := &ListNode{Val:5}
	//l2.Next = &ListNode{Val:6}
	//l2.Next.Next = &ListNode{Val:4}

	//l1 := &ListNode{Val:5}
	//l2 := &ListNode{Val:5}
	l1 := &ListNode{Val:1}
	l2 := &ListNode{Val:9}
	l2.Next = &ListNode{Val:9}
	fmt.Println(l1, l1.Next)
	fmt.Println(l2, l2.Next)
	now := addTwoNumbers(l1, l2)
	fmt.Println(now, now.Next, now.Next.Next)
}

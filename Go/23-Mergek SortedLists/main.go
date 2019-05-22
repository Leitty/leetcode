package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head = &ListNode{0, nil}
	var node *ListNode
	var value int
	node = head
	for !listnil(lists) {
		value, lists = minAndNext(lists)
		node.Next = &ListNode{value, nil}
		node = node.Next
	}
	return head.Next
}

func listnil(lists []*ListNode) bool{
	for _, list := range lists{
		if list != nil {
			return false
		}
	}
	return true
}

func minAndNext(lists []*ListNode) (int, []*ListNode){
	var min = 10000
	for _, list := range lists{
		if list != nil && list.Val < min {
			min = list.Val
		}
	}

	for k, list := range lists{
		if list != nil && list.Val == min {
			lists[k] = lists[k].Next
			break
		}
	}
	return min, lists
}

func main() {
	l1 := &ListNode{1,nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{5, nil}

	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}

	l3 := &ListNode{2, nil}
	l3.Next = &ListNode{6, nil}

	l := []*ListNode{l1, l2, l3}
	t := mergeKLists(l)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}

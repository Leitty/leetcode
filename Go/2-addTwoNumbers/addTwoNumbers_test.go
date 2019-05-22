package main

import (
	"fmt"
	"testing"
)

func GenerateList(val... int) *ListNode{
	head := ListNode{0,nil}
	next := &head

	for _, v := range val {
		next.Next = &ListNode{v,nil}
		next = next.Next
	}

	return  head.Next
}

func EqualList(l1 *ListNode, l2 *ListNode) bool{
	for  p, q := l1 , l2;p != nil || q != nil; {
		if p == nil && q != nil{
			return false
		}
		if p != nil && q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return  true
}

func PrintList(node *ListNode) string{
	s := ""
	p := node
	for p != nil {
		s += fmt.Sprintf("node: %d ", p.Val)
		p = p.Next
	}
	return s
}

func TestAddTwoNumbers(t *testing.T){
	tests := []struct{
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	}{
		{GenerateList(2, 4, 3), GenerateList(5, 6, 4), GenerateList(7, 0, 8)},
		{GenerateList(5), GenerateList(5), GenerateList(0, 1)},
		{GenerateList(1, 9), GenerateList(9), GenerateList(0, 0, 1)},
	}

	for _, tt := range tests {
		if actual := addTwoNumbers(tt.l1, tt.l2); !EqualList(actual, tt.l3){
			t.Errorf("Expect %s, But we get %s", PrintList(tt.l3), PrintList(actual))
		}
	}
}


func BenchmarkAddTwoNumbers(b *testing.B){
	l1 := GenerateList(2, 4, 3)
	l2 := GenerateList(5, 6, 4)
	l3 := GenerateList(7, 0, 8)
	for i :=0 ; i< b.N; i++{
		if actual := addTwoNumbers(l1, l2); !EqualList(actual, l3){
			b.Errorf("Expect %s, But we get %s", PrintList(l3), PrintList(actual))
		}
	}
}



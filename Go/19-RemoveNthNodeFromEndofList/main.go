package main

type ListNode struct {
	Val int
	Next *ListNode
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	var node = head
//	var k = 1
//	for node.Next != nil {
//		k++
//		node = node.Next
//	}
//	if k == 1 {
//		return nil
//	}
//	target := k -n
//	node = head
//	if target == 0 {
//		return head.Next
//	}
//	for ;target != 1;target-- {
//		node = node.Next
//	}
//	node.Next = node.Next.Next
//	return head
//}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var result, first = head, head
	for n>0 {
		first = first.Next
		n--
	}
	if first == nil {
		if n == 1{
			return nil
		}
		return result.Next
	}
	second := head
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return result
}

func main() {

}

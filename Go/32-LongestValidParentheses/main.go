package main

import "fmt"

//使用先进前出队列实现
func longestValidParentheses(s string) int {
	var maxlength,length int
	var queue []int
	for k, v := range s {
		if len(queue)==0{
			queue = append(queue, k)
			continue
		}
		n := queue[len(queue)-1]
		if s[n] == '(' && v == ')' {
			queue = queue[:len(queue)-1]
			if len(queue) != 0 {
				length = k - queue[len(queue)-1]
			} else {
				length = k+1
			}
			if length > maxlength {
				maxlength = length
			}
		} else {
			queue = append(queue, k)
		}
	}
	return maxlength
}


func main() {
	//s := ")()())"
	//s := "(()"
	//s := "(())"
	//s := "(())(())"
	s := "()(())"
	fmt.Println(longestValidParentheses(s))
}

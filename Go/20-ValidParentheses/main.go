package main

import "fmt"

func isValid(s string) bool {
	var queue []rune
	if s == "" {
		return true
	}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			queue = append(queue, v)
		} else {
			len := len(queue)
			if len == 0 {
				queue = append(queue, v)
				continue
			}
			k := queue[len-1]
			if (v == ')' && k == '(') || (v == ']' && k == '[') || (v == '}' && k == '{') {
				queue = queue[:len-1]
				continue
			}
			queue = append(queue, v)
		}
	}
	return len(queue) == 0
}



func main() {
	s := ")"
	fmt.Println(isValid(s))
}

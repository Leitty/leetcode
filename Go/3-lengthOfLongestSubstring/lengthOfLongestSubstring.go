package main

import "fmt"

func lengthOfLongestSubstring(s string) int{
	ch := make(map[rune]int)
	start := 0
	maxLength := 0
	for k, v := range []rune(s) {
		if lastShow, ok := ch[v]; ok && lastShow >= start {
			start = ch[v] +1
		}
		if k-start+1 > maxLength {
			maxLength = k-start+1
		}
		ch[v] = k
	}
	return maxLength
}


func main() {
	fmt.Println(lengthOfLongestSubstring("abcbbaa"))
	fmt.Println(lengthOfLongestSubstring("一二三四二一三"))
}

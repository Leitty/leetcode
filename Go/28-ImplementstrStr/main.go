package main

import "fmt"

func strStr(haystack string, needle string) int {
	l := len(haystack)
	n := len(needle)
	if  n == 0 {
		return 0
	}
	if l == 0 || l < n {
		return -1
	}
	for i:=0;i<l-n+1;i++{
		if haystack[i] == needle[0] {
			j:=0
			for ;j<n;j++{
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == n {
				return i
			}
		}
	}
	return -1
}

// haystack = "hello", needle = "ll"
func main() {
	var haystack = "a"
	var needle = "a"
	fmt.Println(strStr(haystack, needle))
}

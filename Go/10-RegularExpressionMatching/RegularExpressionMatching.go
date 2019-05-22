package main

import "fmt"

func isMatch(s string, p string) bool {
	return ((len(s) <= 0 || len(p) <= 0) && ((len(s) <= 0 && len(p) <= 0) || (len(p) >= 2 && p[1] == '*' && isMatch(s[:], p[2:])))) || (!(len(s) <= 0 || len(p) <= 0)&&(((len(p) > 1 && p[1] == '*') && (isMatch(s[:], p[2:]) || ((p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[:])))) ||  ((p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[1:]))))
}



func main() {
	s := "sssp"
	p := "s*sp"
	fmt.Println(isMatch(s, p))
}

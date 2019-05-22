package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""}
	var lenStr = len(strs)
	var comStr = strs[0]
	for i :=1; i < lenStr; i++ {
		comStr = maxCommonStr(comStr, strs[i])
	}
	return comStr
}

//func maxCommonStr(str string, s string)  string{
//	var lens = len(str)
//	if lens ==1 && len(s)==1 && str == s {
//		return str
//	}
//	for l := lens; l > 1 ; l--{
//		for i := 0 ; i+l < lens+1; i++{
//			substr := str[i:i+l]
//			if strings.Contains(s, substr) {
//				return substr
//			}
//		}
//	}
//	return ""
//}

func maxCommonStr(str string, s string)  string {
	for i := len(str); i >= 0; i--{
		substr := str[0:i]
		if strings.HasPrefix(s, substr){
			return substr
		}
	}
	return ""
}

func main() {
	s :=  []string {"flower","flow","flight"}
	//s :=  []string {"dog","racecar","car"}
	//s := []string {"ca","a"}
	//s := []string {"a"}
	//s := []string {"c","c"}
	//s := []string {"ab","a"}
	//s :=  []string {"c","acc","ccc"}
	fmt.Println(longestCommonPrefix(s))
	//s := "flow"
	//fmt.Println(s[0:4])
}

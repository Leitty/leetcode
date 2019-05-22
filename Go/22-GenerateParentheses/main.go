package main

import (
	"fmt"
)

//func generateParenthesis(n int) []string {
//	if n == 0 {
//		return nil
//	}
//	if n == 1 {
//		return []string{"()"}
//	}
//	var results = []string{}
//	for i := 1 ; i < n ; i++ {
//		substr1 := generateParenthesis(i)
//		substr2 := generateParenthesis(n-i)
//		for _, sub := range stingsMutiStrings(substr1, substr2, i ,n-i) {
//			if len(sub) != 0 {
//				results = append(results, sub)
//			}
//		}
//	}
//	for _ ,sub := range generateParenthesis(n-1) {
//		if len(sub) != 0 {
//			results = append(results, "("+sub+")")
//		}
//	}
//	return results
//}
//
//func removedParenthesis(n int) string {
//	var result string
//	template := "()"
//	for i := 0 ; i< n ;i++ {
//		result = result + template
//	}
//	return result
//}
//
//func stingsMutiStrings(str1 ,str2 []string ,n1, n2 int) []string{
//	if len(str1) == 0 && len(str2) == 0 {
//		return nil
//	}
//	if len(str1) == 0 {
//		return str2
//	}
//	if len(str2) == 0 {
//		return str1
//	}
//	result := []string{}
//	for _, substr1 := range str1 {
//		for _, substr2 := range str2 {
//			if n1 > 1{
//				if substr1 == removedParenthesis(n1)  {
//					continue
//				}
//			}
//			sub := substr1 + substr2
//			result = append(result, sub)
//		}
//	}
//	return result
//}


func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var results = []string{}
	for i := 0 ; i < n ; i++ {
		substr1 := generateParenthesis(i)
		substr2 := generateParenthesis(n-i-1)
		for k, sub := range substr1 {
			substr1[k] = "(" + sub + ")"
		}

		for _, sub := range stingsMutiStrings(substr1, substr2) {
			if len(sub) != 0 {
				results = append(results, sub)
			}
		}
	}
	return results
}

func stingsMutiStrings(str1 ,str2 []string) []string{
	if len(str1) == 0 && len(str2) == 0 {
		return nil
	}
	if len(str1) == 0 {
		return str2
	}
	if len(str2) == 0 {
		return str1
	}
	result := []string{}
	for _, substr1 := range str1 {
		for _, substr2 := range str2 {
			sub := substr1 + substr2
			result = append(result, sub)
		}
	}
	return result
}

func main() {
	n := 4
	fmt.Println(generateParenthesis(n))
}

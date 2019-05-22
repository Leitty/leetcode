package main

import "fmt"

// 方法一: 给string编号,比如: numRows=3时,对应的P数组编号0 1 2 3 2 1 0 1 2 3 2 1 0, 然后再取编号0对应的string,编号1对应的string等等,组成新的字符串
//func convert(s string, numRows int) string {
//	p := make([]int, len(s))
//	k := 0
//	reset := false
//	for i:=0; i<len(s);i++{
//		p[i] = k
//		if !reset {
//			k++
//		} else {
//			k--
//		}
//
//		if k > numRows-1 && !reset {
//			reset = true
//			if k > 1 {
//				k = k-2
//			} else {
//				k = 0
//			}
//		}
//
//		if k < 0 && reset {
//			reset = false
//			k = k+2
//			if k>numRows-1{
//				k = numRows-1
//			}
//		}
//	}
//	fmt.Println(p)
//	result := ""
//	for n:=0;n<numRows;n++{
//		for k,v := range p{
//			if v == n{
//				result = result + string(s[k])
//			}
//		}
//	}
//	return result
//}

// 方法二
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result []byte
	cycleLen := 2*numRows-2

	for i := 0; i < numRows; i++ {
		for j :=i; j < len(s); j = j+cycleLen{
			result= append(result, s[j])
			if i !=0 && i!= numRows-1 && j+cycleLen -2*i <len(s) {
				result = append(result, s[j+cycleLen -2*i])
			}
		}
	}
	return string(result)
}


func main() {
	//s := "PAYPALISHIRING"
	//numRows := 4
	s := "AB"
	numRows := 1
	//s := "ABC"
	//numRows := 1
	//fmt.Println(len(s))
	fmt.Println(convert(s, numRows))
}

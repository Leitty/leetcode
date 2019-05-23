package main

import (
	"fmt"
	"math"
)

//func divide(dividend int, divisor int) int {
//	var i int
//	if divisor == 1 || divisor == -1 {
//		return overflows(dividend/divisor)
//	}
//	if dividend * divisor > 0 {
//		for ;absolute(dividend) >= absolute(divisor);i++{
//			dividend = dividend-divisor
//		}
//		return i
//	} else {
//		for ;absolute(dividend) >= absolute(divisor);i++{
//			dividend = dividend + divisor
//		}
//		return -i
//	}
//}
//
//func absolute(i int) int{
//	if i < 0 {
//		return i * (-1)
//	}
//	return i
//}
//
//func overflows(i int) int{
//	if i > 2147483647 || i < -2147483648 {
//		return 2147483647
//	}
//	return i
//}


func divide(dividend int, divisor int) int {
	var ret, nege int
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}
	if dividend * divisor > 0 {
		nege = 1
	} else {
		nege = -1
	}
	dv1 := absolute(dividend)
	dv2 := absolute(divisor)
	for dv1 >= dv2 {
		i := 1
		for dv1 >= i * dv2 {
			dv1 = dv1 - i * dv2
			ret+=i
			i++
		}
	}
	return nege * ret
}

func absolute(i int) int{
	if i < 0 {
		return i * (-1)
	}
	return i
}

func main() {
	fmt.Println(divide(7, -3))
	//fmt.Println(7>>1)
	//fmt.Println()
}

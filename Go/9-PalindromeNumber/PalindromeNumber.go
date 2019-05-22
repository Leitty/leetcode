package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return  false
	}
	y := reverse(x)
	if x == y {
		return true
	} else {
		return false
	}
}

func reverse(x int) int {
	var target ,m int
	for start:=x; start!=0;  {
		m = start %10
		target = target*10+ m
		start = start/10
	}
	if target < int(math.Pow(-2, 31)) || target > int(math.Pow(2, 31)-1){
		target = 0
	}
	return target
}

func main() {
	//x := 121
	x := 10
	fmt.Println(isPalindrome(x))
}

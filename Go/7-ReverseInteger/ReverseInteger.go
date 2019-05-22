package main

import (
	"fmt"
	"math"
)

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
	x := 1563847412
	fmt.Println(int32(x))
	fmt.Println(reverse(x))
}

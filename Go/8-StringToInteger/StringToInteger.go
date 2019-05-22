package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	max := int(math.Pow(2, 31)-1)
	min := int(math.Pow(-2, 31))
	re := strings.Trim(str," ")
	for k, v := range re{
		if k ==0 && v == 45 {
			continue
		}
		if k ==0 && v == 43 {
			continue
		}
		if v < 48 || v > 57 {
			re = re[:k]
			break
		}
	}
	in, _ := strconv.Atoi(re)
	if in < min{
		in = min
	}
	if in > max{
		in = max
	}
	return in
}

func main() {
	//x :="    -42"
	//x := "words and 987"
	//x := "   -4193 with words"
	//x := "-91283472332"
	x := "+1"
	fmt.Println(myAtoi(x))
}

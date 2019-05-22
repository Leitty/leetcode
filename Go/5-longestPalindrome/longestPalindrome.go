package main

import "fmt"

func longestPalindrome(s string) string{
	// 在s的每个字符间加上#，将s改造成str
	T := "$#"
	for _, v := range []rune(s) {
		T = T + string(v)
		T =T + "#"
	}
	fmt.Println(T)

	// 计算p[]的值
	p := make([]int,len(T))
	right := 0
	mi := 0
	max := 0
	maxi := 0
	for i := 1; i < len(T)-1; i++ {
		if right > i {
			if i+ p[2*mi-i] < right{
				p[i] = p[2*mi-i]
			} else {
				p[i] = right-i
			}
		} else {
			p[i] = 1
		}

		for (i-p[i]) > 0 && (i+p[i]) < len(T)-1 && T[p[i]+i] == T[i-p[i]] {
				p[i]++
		}

		if p[i]+i > right{
			mi = i
			right = p[i]+i
		}

		if p[i] >= max{
			max = p[i]
			maxi = i
		}
	}

	fmt.Printf("Max length: %d, max i: %d\n", max, maxi)
	// 求出最大回环字符串
	return s[(maxi-max)/2:(maxi+max)/2-1]
}

func main() {
	//s := "ahskhdalksjc"
	//s := "abac"
	s := "babad"
	//s := ""
	fmt.Println(longestPalindrome(s))
}

package main

import (
	"fmt"
)

func letterCombinations(digits string)  []string{
	var result, next []string
	for i := 0 ;i< len(digits); i++ {
		list := numToChar(string(digits[i]))
		for _ , l := range list {
			if result == nil {
				next = append(next, l)
			} else {
				for _, v := range result {
					next = append(next, v+l)
					//fmt.Printf("next: %v", next)
				}
			}
		}
		result = next
		next = nil
	}
	return result
}

func numToChar(k string) []string {
	if k == "0" {
		return []string{" "}
	}
	if k == "1" {
		return []string{""}
	}
	if k == "2" {
		return []string{"a","b","c"}
	}
	if k == "3" {
		return []string{"d","e","f"}
	}
	if k == "4" {
		return []string{"g","h","i"}
	}
	if k == "5" {
		return []string{"j","k","l"}
	}
	if k == "6" {
		return []string{"m","n","o"}
	}
	if k == "7" {
		return []string{"p","q","r","s"}
	}
	if k == "8" {
		return []string{"t","u","v"}
	}
	if k == "9" {
		return []string{"w","x","y","z"}
	}
	return nil
}




func main() {
	fmt.Println(letterCombinations("23"))
}

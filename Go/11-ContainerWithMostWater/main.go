package main

import "fmt"

// Brute Force
//func maxArea(height []int) int {
//	var max,square int
//	len := len(height)
//	for i := 1; i< len ; i++ {
//		for j := 0 ; j < i ; j++ {
//			if height[i] > height[j] {
//				square = (i - j ) * height[j]
//			} else {
//				square = (i - j ) * height[i]
//			}
//			if square > max {
//				max = square
//			}
//		}
//	}
//	return max
//}

// Two Pointer Approach
func maxArea(height []int) int {
	left, right := 0 ,len(height)-1
	var max, square int
	for left < right {
		if height[left] > height [right] {
			square = height[right] * (right-left)
			right--
		} else {
			square = height[left] * (right-left)
			left++
		}
		if square > max {
			max = square
		}
	}
	return max
}


func main() {
	//var h [9]int
	h := []int {1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(h))
}

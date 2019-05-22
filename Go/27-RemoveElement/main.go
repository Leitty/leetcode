package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var i int
	for k:=0; k < l; k++ {
		if nums[k] != val {
			nums[i] = nums[k]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums, 2))
}

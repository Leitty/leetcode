package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var  i int
	for k := 1; k <len(nums); k++ {
		if nums[k] != nums[i] {
			i++
			nums[i] = nums[k]
		}
	}

	return i+1
}


func main() {
	nums := []int{0}
	fmt.Println(removeDuplicates(nums))
}

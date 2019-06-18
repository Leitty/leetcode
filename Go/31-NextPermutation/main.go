package main

import (
	"sort"
)

func nextPermutation(nums []int)  {
	ln := len(nums)
	if ln == 1 || ln == 0{
		return
	}
	for l := ln-1; l > 0 ; l--{
		if nums[l-1] < nums[l] {
			sort.Ints(nums[l:])
			for k , v := range nums[l:] {
				if v > nums[l-1] {
					nums[k+l] = nums[l-1]
					nums[l-1] = v
					return
				}
			}
		}
	}
	sort.Ints(nums)
}



func main() {

}

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	length := len(nums)
	if length == 0 || length < 4 {
		for _, v := range  nums {
			result +=v
		}
		return result
	}
	sort.Ints(nums)
	result = nums[0]+nums[1]+nums[2]
	for i := 0; i< length - 2; i++ {
		want := target - nums[i]
		for left, right := i+1, length-1 ; left < right; {
			sum := nums[left] + nums[right]
			if  sum == want {
				return target
			}
			if sum < want {
				left++
			} else if sum > want {
				right--
			}
			current := sum + nums[i]
			if math.Abs(float64(current-target)) < math.Abs(float64(result-target)){
				result = current
			}
		}
	}
	return result
}

func main() {
	n := []int{0}
	fmt.Println(threeSumClosest(n, 1))
}

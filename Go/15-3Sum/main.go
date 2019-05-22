package main

import (
	"fmt"
	"sort"
)

//func threeSum(nums []int) [][]int {
//	var result [][]int
//	if len(nums) < 3 {
//		return result
//	}
//	sort.Ints(nums)
//	lenn := len(nums)-1
//	for i := 0 ; i< lenn-1; i++ {
//		if i > 1 && nums[i] == nums[i-1] {
//			continue
//		}
//		var left, right = i+1, lenn
//		target := -nums[i]
//		for left < right  {
//			sum := nums[left] + nums[right]
//			if sum > target {
//				right--
//			} else if sum < target{
//				left++
//			} else if sum == target {
//				result = append(result, []int{nums[i], nums[left], nums[right]})
//				left++
//				right--
//				for left < right && nums[left] == nums[left-1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right+1] {
//					right--
//				}
//			}
//
//		}
//	}
//	return result
//}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	n := len(nums)
	if n == 0 || n < 3 {
		return results
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}

func main() {
	//nums := []int {-1, 0, 1, 2, -1, -4}
	nums := []int {0,0,0,0}
	//sort.Ints(nums)
	fmt.Println(threeSum(nums))
}

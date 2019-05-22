package main

import (
	"fmt"
	"sort"
)

//func fourSum(nums []int, target int) [][]int {
//	results := [][]int{}
//	n := len(nums)
//	if n < 4 {
//		return results
//	}
//	sort.Ints(nums)
//	for k:=0; k<n-3; k++ {
//		if k>0 && nums[k] == nums[k-1] {
//			continue
//		}
//		threes := threeSum(nums[k+1:], target-nums[k])
//		if len(threes) == 0 {
//			continue
//		}
//		for _, i := range threes{
//			t := append(i, nums[k])
//			results = append(results, t)
//		}
//	}
//	return results
//}
//
//func threeSum(nums []int, want int) [][]int {
//	results := [][]int{}
//	n := len(nums)
//	if n == 0 || n < 3 {
//		return results
//	}
//	sort.Ints(nums)
//	for i := 0; i < n-2; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		target := want-nums[i]
//		left := i + 1
//		right := n - 1
//		for left < right {
//			sum := nums[left] + nums[right]
//			if sum == target {
//				results = append(results, []int{nums[left], nums[right], nums[i]})
//				left++
//				right--
//				for left < right && nums[left] == nums[left-1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right+1] {
//					right--
//				}
//			} else if sum > target {
//				right--
//			} else if sum < target {
//				left++
//			}
//		}
//	}
//	return results
//}


func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return NSum(nums, target, 4)
}

func NSum(nums []int, target, N int) [][]int {
	results := [][]int{}
	if len(nums) ==0 {
		return results
	}
	if N == 0 {
		return results
	}
	if N == 1 {
		for _, v := range nums{
			if v == target {
				return [][]int{{v}}
			}
		}
	}
	if N > len(nums) {
		return results
	}

	n1sums := NSum(nums[1:], target, N)
	if len(n1sums) != 0 {
		for _, v := range n1sums {
			results = append(results, v)
		}
	}

	n2sums := NSum(nums[1:], target-nums[0], N-1)
	if len(n2sums) != 0 {
		for _, v := range n2sums{
			result := append(v, nums[0])
			results = append(results, result)
		}
	}

	return results
}


func main() {
	//nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{-1,-5,-5,-3,2,5,0,4}
	fmt.Println(fourSum(nums,-7))
	//fmt.Println([][]int{{1}})
}

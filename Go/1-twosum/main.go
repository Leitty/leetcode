package main

import (
	"fmt"
)

//func twoSum(nums []int, target int) []int {
//	for i := range nums {
//		another := target - nums[i]
//		for  j := range  nums {
//			if nums[j] == another && j != i {
//				return  []int{i,j}
//			}
//		}
//	}
//	return nil
//}

//func twoSum(nums []int, target int) []int {
//	//var NumMap = make([]int,0xffff)
//
//}

func twoSum(nums []int, target int) []int {
	//sort.Ints(nums)
	NumMap := make(map[int]int)
	for k, v := range nums {
		NumMap[v] = k
	}

	for k, v := range nums {
		if m, ok := NumMap[target-v]; ok && m != k {
			return []int{k, m}
		}
	}
	return nil
}

//func twoSum(nums []int, target int) []int {
//	var x []int
//	num:=len(nums)
//	xx:=make(map[int]int)//用一个map存下每个元素值，如果（target - 一个元素值）能在这个map里找到，直接就得到了满足题意的元素组合
//	for i:=0;i<num;i++{
//		xx[nums[i]]=i
//	}
//	for i:=0;i<num;i++{
//		if thenum,ok:=xx[target-nums[i]];ok {
//			if thenum==i{
//				continue
//			}
//			x = append(x, i)
//			x = append(x, thenum)
//			return x
//		}
//	}
//	return x
//}

func main() {
	nums := []int{3,2,4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

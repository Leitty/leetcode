package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct{
		nums []int
		target int
		result []int
	}{
		{[]int{1,3,11,18},4, []int{0,1}},
		{[]int{3, 3},6, []int{0, 1}},
		{[]int{3,2,4}, 6, []int{1,2}},
		{[]int{2, 5},6,[]int{}},
		//{[]int{3, 3, 3}, 6, []int{0, 1}},
	}

	for _,tt := range tests {
		actual := twoSum(tt.nums, tt.target)
		//for k,v := range actual {
		//	if tt.result[k] != v {
		//		t.Errorf("twoSum(%v, %d) result in %v, the actual result is %v",
		//			tt.nums, tt.target, actual, tt.result )
		//	}
		//}
		if !equal(actual,tt.result){
			t.Errorf("twoSum(%v, %d) result in %v, the actual result is %v",
						tt.nums, tt.target, actual, tt.result)
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

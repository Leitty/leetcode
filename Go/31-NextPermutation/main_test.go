package main

import "testing"

func TestNextPermutation(t *testing.T) {
	var tests = []struct{
		nums []int
		target []int
	} {
		{[]int{1,2,3},[]int{1,3,2}},
		{[]int{3,2,1},[]int{1,2,3}},
		{[]int{1,1,5},[]int{1,5,1}},
		{[]int{2,4,3,1},[]int{3,1,2,4}},
		{[]int{5,2,4,3,1},[]int{5,3,1,2,4}},
		{[]int{2,4,5,3,1},[]int{2,5,1,3,4}},
	}

	for _, test := range tests {
		nextPermutation(test.nums)
		for k, v := range test.nums {
			if v != test.target[k] {
				t.Errorf("Wanted is %v, but get %v\n", test.target, test.nums)
				break
			}
		}
	}
}
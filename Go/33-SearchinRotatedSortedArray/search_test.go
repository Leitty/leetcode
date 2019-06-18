package _3_SearchinRotatedSortedArray

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct{
		nums []int
		target int
		want int
	} {
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{4,5,6,7,0,1,2}, 3, -1},
		{[]int{4,5,6,7,0,1,2}, 2, 6},
		{[]int{5,1,3}, 3, 2},
		{[]int{3,1}, 3, 0},
	}

	for _, test := range tests {
		result := Search(test.nums, test.target)
		if result != test.want {
			t.Errorf("Result is %v, but wanted is %v\n", result, test.want)
		}
	}
}


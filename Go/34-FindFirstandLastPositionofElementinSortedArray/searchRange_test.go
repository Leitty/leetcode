package _4_FindFirstandLastPositionofElementinSortedArray

import "testing"

func TestSearchRange(t *testing.T) {
	var tests = []struct{
		Nums []int
		Target int
		Want []int
	}{
		{[]int{}, 8,[]int{-1, -1}},
		{[]int{}, 0,[]int{-1, -1}},
		{[]int{1}, 8,[]int{-1, -1}},
		{[]int{1}, 1,[]int{0, 0}},
		{[]int{1,1}, 1,[]int{0,1}},
		{[]int{1,2}, 8,[]int{-1,-1}},
		{[]int{5,7,7,8,8,10}, 8,[]int{3,4}},
		{[]int{5,7,7,8,8,10}, 6,[]int{-1, -1}},
		{[]int{1,2,3,3,4,5,6,7,8,8,9,10}, 3,[]int{2, 3}},
	}

	for _, test := range tests {
		result := SearchRange(test.Nums, test.Target)
		t.Logf("Result is: %v\n", result)
		for k := range result {
			if result[k] != test.Want[k] {
				t.Errorf("Want is %v, but result is %v\n", result[k], test.Want[k])
			}
		}
	}

}



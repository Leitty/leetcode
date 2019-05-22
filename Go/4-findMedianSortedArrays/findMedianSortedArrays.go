package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//var arr = make([]int, len(nums1)+len(nums2))
	var arr []int
	m, n  := 0, 0
	lm , ln := len(nums1), len(nums2)
	for {
		if m == lm && n == ln{
			break
		}

		if m == lm {
			for n != ln {
				arr = append(arr, nums2[n])
				n++
			}
			break
		}

		if n == ln {
			for m != lm {
				arr = append(arr, nums1[m])
				m++
			}
			break
		}

		if nums1[m] <= nums2[n] {
			arr = append(arr, nums1[m])
			m++
		}else {
				arr = append(arr, nums2[n])
				n++
		}
	}
	if ( (lm+ln)%2 == 0 ) {
		re := float64(arr[(lm+ln)/2-1]+arr[(lm+ln)/2])/2
		return float64(re)
	}else {
		re := arr[((lm+ln)-1)/2]
		return  float64(re)
	}
}

func main() {
	//nums1 := []int{1, 3}
	//nums2 := []int{2}

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	result := findMedianSortedArrays(nums1,nums2)
	fmt.Printf("The result : %f\n", result)
}


// 1 3 3 4 5
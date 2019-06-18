package _3_SearchinRotatedSortedArray

//func Search(nums []int, target int) int {
//	var i int
//	ln := len(nums)
//	if len(nums) == 0 {
//		return -1
//	}
//
//	if len(nums) == 1 && nums[0] != target {
//		return -1
//	}
//
//	if target == nums[ln/2] {
//		return ln/2
//	}
//	if nums[0] <= nums[ln/2] && target >= nums[0] && target <= nums[ln/2] ||
//		nums[0] > nums[ln/2] && (target >= nums[0] || target <= nums[ln/2]) {
//		i = Search(nums[:ln/2], target)
//		if i == -1 {
//			return -1
//		} else {
//			return i
//		}
//	} else {
//		i = Search(nums[(ln/2+1):], target)
//		if i == -1 {
//			return -1
//		} else {
//			return ln/2 + 1 + i
//		}
//	}
//}

func Search(nums []int, target int) int {
	ln := len(nums)
	if ln == 0 {
		return -1
	}
	start, end := 0, ln-1
	for start < end {
		mid := (start + end)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] && target >= nums[0] && target <= nums[mid] ||
			nums[0] > nums[mid] && (target >= nums[0] || target <= nums[mid]) {
			end = mid
		} else {
			start = mid+1
		}
	}

	if nums[start] == target {
		return start
	}
	return -1
}
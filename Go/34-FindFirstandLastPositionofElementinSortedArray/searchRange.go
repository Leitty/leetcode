package _4_FindFirstandLastPositionofElementinSortedArray


func SearchRange(nums []int, target int) []int {
	ln := len(nums)
	if ln == 0 {
		return []int{-1,-1}
	}
	if ln == 1 {
		if nums[0] == target {
			return []int{0,0}
		} else {
			return []int{-1,-1}
		}
	}
	var start, end = 0, ln-1
	for start < end {
		if nums[start] == target && nums[end] == target {
			break
		}
		mid := (start+end)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] == target {
			if nums[start] != target {
				lstart := (mid-start)/2
				for ; lstart > 1;lstart = lstart/2{
					if nums[start+lstart] < target {
						break
					}
				}
				if lstart < 1 {
					lstart = 1
				}
				start = start + lstart
			}

			if nums[end] != target {
				lend := (end-mid)/2
				for ; lend > 1;lend = lend/2{
					if nums[end-lend] > target {
						break
					}
				}
				if lend < 1 {
					lend = 1
				}
				end = end - lend
			}
		}
	}
	if nums[start] != target {
		start = -1
		end = -1
	}
	return []int{start, end}
}
package problems

func binarySearch(lo int, hi int, nums []int, target int) int {
	r := -1
	if lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}

		// if target < nums[mid] {
		r = binarySearch(lo, mid-1, nums, target)
		if r == -1 {
			r = binarySearch(mid+1, hi, nums, target)
		}
	}

	if r != -1 {
		return r
	}
	return -1
}

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	r := binarySearch(lo, hi, nums, target)
	return r
}

func search2(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	r := -1

	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}

		if mid > 0 && nums[lo] < nums[mid-1] {
			if nums[lo] <= target && target <= nums[mid-1] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if mid < len(nums)-1 && nums[mid+1] <= target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return r
}

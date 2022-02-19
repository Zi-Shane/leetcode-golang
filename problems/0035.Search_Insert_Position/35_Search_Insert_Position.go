package problems

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for (end - start) > 0 {
		// search
		p := (end + start) / 2
		if target > nums[p] {
			start = p + 1
		} else if target < nums[p] {
			end = p
		} else {
			return p
		}
	}
	if target > nums[start] {
		return start + 1
	} else {
		return start
	}
}

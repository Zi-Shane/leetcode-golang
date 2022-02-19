package problems

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	// <= i
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i, v := range nums {
		if v == 0 {
			res[i] = 0
			continue
		}
		res[i] = count[v-1]
	}

	return res
}

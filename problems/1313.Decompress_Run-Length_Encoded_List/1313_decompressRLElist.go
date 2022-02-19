package problems

func decompressRLElist(nums []int) []int {
	length := 0
	for i := 0; i < len(nums); i += 2 {
		length += nums[i]
	}

	result := make([]int, length)
	index := 0
	for i := 0; i < len(nums); i += 2 {
		for count := 0; count < nums[i]; count++ {
			result[index] = nums[i+1]
			index++
		}
	}

	return result
}

func decompressRLElist2(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums)/2; i++ {
		for count := 0; count < nums[2*i]; count++ {
			result = append(result, nums[2*i+1])
		}
	}

	return result
}

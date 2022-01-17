package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := 0
	res := make([]bool, len(candies))

	// find maximun number
	for _, v := range candies {
		if v > maxNum {
			maxNum = v
		}
	}

	// compare
	for i, v := range candies {
		res[i] = maxNum-v <= extraCandies
	}

	return res
}

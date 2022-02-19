package problems

func numIdenticalPairs(nums []int) int {
	res := 0
	counts := make(map[int]int)
	for _, v := range nums {
		res += counts[v]
		counts[v]++
	}
	// {1: 3, 2: 1, 3: 2}

	return res
}

package problems

func combine(n int, k int) [][]int {
	res := [][]int{}
	// cur := []int{}

	var helper func(start int, cur []int)
	helper = func(start int, cur []int) {
		if len(cur) == k {
			dst := make([]int, k)
			copy(dst, cur)
			res = append(res, cur)
		} else {
			for i := start; i <= n; i++ {
				cur = append(cur, i)
				helper(i+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	helper(1, []int{})

	return res
}

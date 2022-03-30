package problems

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	for _, v := range nums {
		lenOfRes := len(res)
		for i := 0; i < lenOfRes; i++ {
			dst := make([]int, len(res[i]))
			copy(dst, res[i])
			// dst := append(dst, v)
			res = append(res, append(dst, v))
		}
	}

	return res
}

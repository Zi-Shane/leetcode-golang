package problems

import "sort"

func back40(start int, candidates []int, res *[][]int, cur *[]int, sum int, target int) {
	if sum == target {
		dst := make([]int, len(*cur))
		copy(dst, *cur)
		*res = append(*res, dst)
	}
	for i := start; i < len(candidates); i++ {
		if sum < target {
			*cur = append(*cur, candidates[i])
			sum += (*cur)[len(*cur)-1]
			back40(i+1, candidates, res, cur, sum, target)
			if sum > target {
				sum -= (*cur)[len(*cur)-1]
				*cur = (*cur)[:len(*cur)-1]
				break
			}
			sum -= (*cur)[len(*cur)-1]
			*cur = (*cur)[:len(*cur)-1]
		}

		// is sorted, so break loop
		if sum > target {
			break
		}

		// jump to next index, if same "candidates[i]" appear
		for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
			i++
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}

	sort.Ints(candidates)
	back40(0, candidates, &res, &[]int{}, 0, target)

	return res
}

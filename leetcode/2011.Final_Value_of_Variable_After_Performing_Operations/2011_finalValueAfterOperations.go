package leetcode

func finalValueAfterOperations(operations []string) int {
	op := map[string]int{"X++": 1, "++X": 1, "X--": -1, "--X": -1}
	res := 0
	for _, v := range operations {
		num, ok := op[v]
		if ok {
			res += num
		}
	}

	return res
}

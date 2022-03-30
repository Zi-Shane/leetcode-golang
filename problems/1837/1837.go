package problems

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	umaMap := make(map[int]map[int]bool)
	ans := make([]int, k)

	for _, v := range logs {
		id := v[0]
		t := v[1]
		_, exist := umaMap[id]
		if !exist {
			umaMap[id] = make(map[int]bool)
		}
		umaMap[id][t] = true
	}

	for _, v := range umaMap {
		ans[len(v)-1]++
	}

	return ans
}

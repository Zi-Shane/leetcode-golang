package problems

func processQueries(queries []int, m int) []int {
	res := make([]int, len(queries))
	P := make(map[int]int)
	for i := 0; i < m; i++ {
		P[i+1] = i
	}
	for i, v := range queries {
		res[i] = P[v]
		for j, pos := range P {
			if pos < P[v] {
				P[j]++
			}
		}
		P[v] = 0
	}

	return res
}

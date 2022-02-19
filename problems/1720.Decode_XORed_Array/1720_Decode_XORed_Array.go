package problems

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 0; i < len(encoded); i++ {
		res[i+1] = encoded[i] ^ res[i]
	}

	return res
}

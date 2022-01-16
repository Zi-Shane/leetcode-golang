package leetcode

func numJewelsInStones(jewels string, stones string) int {
	res := 0
	counts := make(map[byte]int)
	for _, v := range stones {
		counts[byte(v)]++
	}

	for _, v := range jewels {
		res += counts[byte(v)]
	}

	return res
}

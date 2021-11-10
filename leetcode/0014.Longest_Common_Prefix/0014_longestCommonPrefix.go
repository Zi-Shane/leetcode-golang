package leetcode

func longestCommonPrefix(strs []string) string {
	var ans []byte

	if len(strs) == 1 {
		return strs[0]
	}

	for i, c := range strs[0] {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || byte(c) != strs[j][i] {
				return string(ans)
			}
		}
		ans = append(ans, byte(c))
	}

	return string(ans)
}

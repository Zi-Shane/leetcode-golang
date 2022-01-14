package leetcode

import "strings"

func mostWordsFound(sentences []string) int {
	res := 0
	for _, s := range sentences {
		LenS := strings.Count(s, " ") + 1
		if LenS > res {
			res = LenS
		}
	}

	return res
}

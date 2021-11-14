package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		hashMap[string(bs)] = append(hashMap[string(bs)], s)
	}
	var ans [][]string
	for _, v := range hashMap {
		ans = append(ans, v)
	}

	return ans
}

func groupAnagrams2(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, s := range strs {
		var tmp [26]int
		for i := 0; i < len(s); i++ {
			tmp[s[i]-'a']++
		}
		hashMap[tmp] = append(hashMap[tmp], s)
	}
	ans := [][]string{}
	for _, v := range hashMap {
		ans = append(ans, v)
	}

	return ans
}

package leetcode

func lengthOfLongestSubstring(s string) int {
	left := 0
	max_len := 0
	hash_map := make(map[byte]int)
	for right, c := range s {
		if val, ok := hash_map[byte(c)]; !ok || val < left {
			hash_map[byte(c)] = right
		} else {
			tmp := right - left
			if tmp > max_len {
				max_len = tmp
			}
			left = hash_map[byte(c)] + 1
			hash_map[byte(c)] = right
		}
		tmp := right - left + 1
		if tmp > max_len {
			max_len = tmp
		}
	}
	return max_len
}

func lengthOfLongestSubstring2(s string) int {
	left := 0
	max_len := 0
	hash_map := make(map[byte]int)
	for right, c := range s {
		if val, ok := hash_map[byte(c)]; ok {
			if (val + 1) > left {
				left = val + 1
			}
		}
		hash_map[byte(c)] = right
		if (right - left + 1) > max_len {
			max_len = right - left + 1
		}
	}
	return max_len
}

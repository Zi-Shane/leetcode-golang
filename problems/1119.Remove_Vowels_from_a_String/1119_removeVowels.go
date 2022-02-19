package problems

import "regexp"

func removeVowels(input string) string {
	output := ""
	vowels := map[byte]byte{'a': 'a', 'e': 'e', 'i': 'i', 'o': 'o', 'u': 'u'}
	for i := 0; i < len(input); i++ {
		_, isVowel := vowels[input[i]]
		if isVowel {
			continue
		}
		output += string(input[i])
	}
	return output
}

func removeVowels2(S string) string {
	reg := regexp.MustCompile(`(a|e|i|o|u)`)
	return reg.ReplaceAllString(S, "")
}

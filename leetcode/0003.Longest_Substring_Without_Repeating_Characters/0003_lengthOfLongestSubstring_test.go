package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem3(t *testing.T) {
	// input := "abcabcbbbb"
	// input := " "
	// input := "tmmzuxt"
	input := "pwwkew"
	var ans int
	ans = lengthOfLongestSubstring(input)
	fmt.Println("lengthOfLongestSubstring(): ", ans)
	ans = lengthOfLongestSubstring2(input)
	fmt.Println("lengthOfLongestSubstring2(): ", ans)

}

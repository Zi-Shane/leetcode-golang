package problems

import (
	"fmt"
	"testing"
)

func Test_Problem14(t *testing.T) {
	strs := [][]string{{"flower", "flow", "flight"}, {"a"}, {"ab", "a"}, {"flower", "flower", "flower"}, {""}}
	var ans string
	for i := 0; i < len(strs); i++ {
		ans = longestCommonPrefix(strs[i])
		fmt.Println(ans)
	}
}

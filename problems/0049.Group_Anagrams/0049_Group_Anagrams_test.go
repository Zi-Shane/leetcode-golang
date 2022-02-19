package problems

import (
	"fmt"
	"testing"
)

func Test_Problem49(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
}

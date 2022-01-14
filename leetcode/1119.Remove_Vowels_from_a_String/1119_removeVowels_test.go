package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveVowels(t *testing.T) {
	testCases := []struct {
		para string
		ans  string
	}{
		{
			para: "leetcodeisacommunityforcoders",
			ans:  "ltcdscmmntyfrcdrs",
		},
		{
			para: "aeiou",
			ans:  "",
		},
	}
	for _, tC := range testCases {
		a := removeVowels(tC.para)
		if a == tC.ans {
			fmt.Printf("test case (%v): PASS\n", tC.para)
			t.Log("Correct Answer")
		} else {
			t.Errorf("Wrong Answer: \noutput: %v \nanswer: %v", a, tC.ans)
		}
	}
}

func BenchmarkRemoveVowels01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeVowels("leetcodeisacommunityforcoders")
	}
}

func BenchmarkRemoveVowels02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeVowels2("leetcodeisacommunityforcoders")
	}
}

package leetcode

import "testing"

func TestMostWordsFound(t *testing.T) {
	testCases := []struct {
		sentences []string
		ans       int
	}{
		{
			sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			ans:       6,
		},
	}
	for _, tC := range testCases {
		res := mostWordsFound(tC.sentences)
		if res != tC.ans {
			t.Errorf("Wrong Answer")
		}
	}
}

package leetcode

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	testCases := []struct {
		operations []string
		ans        int
	}{
		{
			operations: []string{"--X", "X++", "X++"},
			ans:        1,
		},
		{
			operations: []string{"++X", "++X", "X++"},
			ans:        3,
		},
		{
			operations: []string{"X++", "++X", "--X", "X--"},
			ans:        0,
		},
	}
	for i, tC := range testCases {
		res := finalValueAfterOperations(tC.operations)
		if res != tC.ans {
			t.Errorf("Case %v: Wrong Answer!!", i)
		}
	}
}

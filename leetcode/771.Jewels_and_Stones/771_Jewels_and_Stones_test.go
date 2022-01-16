package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewels string
		stones string
		ans    int
	}{
		{
			jewels: "aA",
			stones: "aAAbbbb",
			ans:    3,
		},
		{
			jewels: "z",
			stones: "ZZ",
			ans:    0,
		},
	}
	for i, tC := range testCases {
		res := numJewelsInStones(tC.jewels, tC.stones)
		if res != tC.ans {
			t.Errorf("TestCase %v is wrong!! res=%v, ans=%v", i, res, tC.ans)
		}
	}
}

package problems

import "testing"

func TestDecode(t *testing.T) {
	testCases := []struct {
		encoded []int
		first   int
		ans     []int
	}{
		{
			encoded: []int{1, 2, 3},
			first:   1,
			ans:     []int{1, 0, 2, 1},
		},
		{
			encoded: []int{6, 2, 7, 3},
			first:   4,
			ans:     []int{4, 2, 0, 7, 4},
		},
	}
	for _, tC := range testCases {
		res := decode(tC.encoded, tC.first)
		for i, v := range res {
			if v != tC.ans[i] {
				t.Errorf("TestCase%v is wrong!! res=%v, ans=%v", i, res, tC.ans)
			}
		}
	}
}

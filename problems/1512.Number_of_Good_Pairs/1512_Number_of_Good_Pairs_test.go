package problems

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []struct {
		input []int
		ans   int
	}{
		{
			input: []int{1, 2, 3, 1, 1, 3},
			ans:   4,
		},
		{
			input: []int{1, 1, 1, 1},
			ans:   6,
		},
	}
	for i, tC := range testCases {
		res := numIdenticalPairs(tC.input)
		if tC.ans != res {
			t.Errorf("TestCase %v is wrong!! res=%v, ans=%v", i, res, tC.ans)
		}
	}

}

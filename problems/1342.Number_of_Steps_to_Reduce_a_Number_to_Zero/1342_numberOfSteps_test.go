package problems

import "testing"

func TestNumberOfSteps(t *testing.T) {
	testCases := []struct {
		num int
		ans int
	}{
		{
			num: 14,
			ans: 6,
		},
	}
	for i, tC := range testCases {
		res := numberOfSteps(tC.num)
		if res != tC.ans {
			t.Errorf("TestCase%v is wrong!! res=%v, ans=%v", i, res, tC.ans)
		}
	}
}

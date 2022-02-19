package problems

import "testing"

func TestMinPartitions(t *testing.T) {
	testCases := []struct {
		n   string
		ans int
	}{
		{
			n:   "32",
			ans: 3,
		},
		{
			n:   "82734",
			ans: 8,
		},
		{
			n:   "27346209830709182346",
			ans: 9,
		},
	}
	for i, tC := range testCases {
		res := minPartitions(tC.n)
		if res != tC.ans {
			t.Errorf("TestCase %v is wrong! res=%v", i, res)
		}
	}
}

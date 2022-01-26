package leetcode

import "testing"

func TestCountPoints(t *testing.T) {
	testCases := []struct {
		points  [][]int
		queries [][]int
		ans     []int
	}{
		{
			points:  [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			queries: [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			ans:     []int{3, 2, 2},
		},
	}
	for _, tC := range testCases {
		res := countPoints(tC.points, tC.queries)
		for i, v := range tC.ans {
			if v != res[i] {
				t.Errorf("TestCase %v is wrong! res=%v, ans=%v", i, res, tC.ans)
			}
		}
	}
}

package problems

import (
	"fmt"
	"testing"
)

func TestUpdateSubrectangle(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestGetValue(t *testing.T) {
	testCases := []struct {
		row int
		col int
		val int
	}{
		{
			row: 0,
			col: 2,
			val: 1,
		},
	}
	subrec := Constructor([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	for _, tC := range testCases {
		res := subrec.GetValue(tC.row, tC.col)
		if res != tC.val {
			t.Errorf("Wrong value!! (%v, %v) get %v != %v", tC.row, tC.col, res, tC.val)
		} else {
			fmt.Printf("(%v, %v) get %v\n", tC.row, tC.val, tC.val)
		}
	}
}

package problems

import (
	"fmt"
	"testing"
)

func TestDiagonalSort(t *testing.T) {
	testCases := []struct {
		mat [][]int
	}{
		{
			mat: [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}},
		},
		{
			mat: [][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}},
		},
	}
	for _, tC := range testCases {
		diagonalSort(tC.mat)

		for i := 0; i < len(tC.mat); i++ {
			for j := 0; j < len(tC.mat[0]); j++ {
				fmt.Print(tC.mat[i][j], " ")
			}
			fmt.Println("")
		}
	}
}

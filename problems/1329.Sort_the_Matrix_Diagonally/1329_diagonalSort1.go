package problems

import "sort"

func diagonalSort1(mat [][]int) [][]int {
	height := len(mat)
	width := len(mat[0])
	for c := 0; c < width; c++ {
		// mat[0, c]
		sort.Sort(&diag{mat, 0, c})
	}

	for r := 1; r < height; r++ {
		// mat[r, 0]
		sort.Sort(diag{mat, r, 0})
	}

	return mat
}

// (0,3),
// (0,2), (1,3),
// (0,1), (1,2), (2,3)
// (0,0), (1,1), (2,2)
// (1,0), (2,1)
// (2,0)

type diag struct {
	a [][]int
	r int
	c int
}

func (d diag) Len() int {
	dc := len(d.a[0]) - d.c
	dr := len(d.a) - d.r
	if dc < dr {
		return dc
	} else {
		return dr
	}
}

func (d diag) Less(i, j int) bool {
	return d.a[d.r+i][d.c+i] < d.a[d.r+j][d.c+j]
}

func (d diag) Swap(i, j int) {
	d.a[d.r+i][d.c+i], d.a[d.r+j][d.c+j] = d.a[d.r+j][d.c+j], d.a[d.r+i][d.c+i]
}

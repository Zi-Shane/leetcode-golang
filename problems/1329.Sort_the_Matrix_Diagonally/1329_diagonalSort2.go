package problems

import "sort"

func diagonalSort(mat [][]int) [][]int {
	height := len(mat)
	width := len(mat[0])
	diagMap := make(map[int][]int)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			diagMap[i-j] = append(diagMap[i-j], mat[i][j])
		}
	}

	for _, v := range diagMap {
		sort.Ints(v)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			mat[i][j] = diagMap[i-j][0]
			diagMap[i-j] = diagMap[i-j][1:]
		}
	}

	return mat
}

// (0,3),
// (0,2), (1,3),
// (0,1), (1,2), (2,3)
// (0,0), (1,1), (2,2)
// (1,0), (2,1)
// (2,0)

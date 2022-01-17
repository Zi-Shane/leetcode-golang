package leetcode

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sq.rectangle[i][j] = newValue
		}
	}
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	return sq.rectangle[row][col]
}

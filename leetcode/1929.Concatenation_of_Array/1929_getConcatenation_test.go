package leetcode

import (
	"fmt"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{
			input: []int{1, 2, 1},
		},
		{
			input: []int{1, 3, 2, 1},
		},
	}
	for _, tC := range testCases {
		res := getConcatenation(tC.input)
		fmt.Println(res)
	}
}

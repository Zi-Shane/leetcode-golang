package problems

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		arr [][]int
	}{
		{
			arr: [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}},
		},
	}
	for _, tC := range testCases {
		fmt.Println(findingUsersActiveMinutes(tC.arr, 5))
	}
}

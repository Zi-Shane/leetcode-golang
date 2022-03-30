package problems

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		s string
	}{
		{
			s: "A",
		},
	}
	for _, tC := range testCases {
		fmt.Println(convert(tC.s, 1))
	}
}

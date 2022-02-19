package problems

import (
	"fmt"
	"testing"
)

func TestProcessQueries(t *testing.T) {
	q := []int{3, 1, 2, 1}
	m := 5
	fmt.Println(processQueries(q, m))
}

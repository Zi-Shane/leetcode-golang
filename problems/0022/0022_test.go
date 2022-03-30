package problems

import (
	"fmt"
	"testing"
)

func TestGenerateAll(t *testing.T) {
	res := generateParenthesis3(3)

	fmt.Println(res)
}

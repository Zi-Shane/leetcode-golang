package problems

import (
	"fmt"
	"testing"
)

func Test_Problem206(t *testing.T) {
	input := []byte("Hello")
	reverseString(input)
	fmt.Println(string(input))
}

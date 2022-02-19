package problems

import (
	"fmt"
	"testing"
)

type question412 struct {
	para412
	ans412
}

// para 是参数
// one 代表第一个参数
type para412 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans412 struct {
	one []string
}

func Test_Problem412(t *testing.T) {

	qs := []question412{

		{
			para412{15},
			ans412{[]string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 412------------------------\n")

	for _, q := range qs {
		_, p := q.ans412, q.para412
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, fizzBuzz(p.one))
	}
	fmt.Printf("\n\n\n")
}

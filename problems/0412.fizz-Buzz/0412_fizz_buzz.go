package problems

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	words := map[int]string{3: "Fizz", 5: "Buzz"}
	divs := []int{3, 5}

	for i := 0; i < n; i++ {
		s := ""
		for _, val := range divs {
			if ((i + 1) % val) == 0 {
				s += words[val]
			}
		}
		if s == "" {
			s = strconv.Itoa(i + 1)
		}
		ans[i] = s
	}

	return ans
}

// func main() {
// 	fmt.Print(fizzBuzz(15))
// }

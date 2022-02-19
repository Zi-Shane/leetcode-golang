package problems

func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			num -= 1
		} else {
			num /= 2
		}
		res++
	}

	return res
}

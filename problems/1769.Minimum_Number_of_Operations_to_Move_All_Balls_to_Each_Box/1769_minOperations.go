package problems

func minOperations(boxes string) []int {
	res := make([]int, len(boxes))
	cnt := 0
	ops := 0
	for i := 0; i < len(boxes); i++ {
		res[i] += ops
		cnt += int(boxes[i] - '0')
		ops += cnt
	}
	cnt = 0
	ops = 0
	for j := len(boxes) - 1; j >= 0; j-- {
		res[j] += ops
		cnt += int(boxes[j] - '0')
		ops += cnt
	}

	return res
}

func minOperations2(boxes string) []int {
	res := make([]int, len(boxes))
	cnt := 0
	ops := 0
	for i := 0; i < len(boxes); i++ {
		res[i] += cnt*i - ops
		if boxes[i] == '1' {
			cnt++
			ops += i
		}
	}
	cnt = 0
	ops = 0
	for j := len(boxes) - 1; j >= 0; j-- {
		res[j] += ops - cnt*j
		if boxes[j] == '1' {
			cnt++
			ops += j
		}
	}

	return res
}

func minOperations1(boxes string) []int {
	pos := []int{}
	for i, ch := range boxes {
		if ch == '1' {
			pos = append(pos, i)
		}
	}

	res := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		count := 0
		for j := 0; j < len(pos); j++ {
			diff := pos[j] - i
			if diff < 0 {
				count += -diff
			} else {
				count += diff
			}
		}
		res[i] = count
	}

	return res
}

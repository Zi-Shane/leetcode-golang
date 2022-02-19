package problems

func minPartitions(n string) int {
	maxNum := 0
	for i := 0; i < len(n); i++ {
		byteInt := int(n[i] - '0')
		if byteInt > maxNum {
			maxNum = byteInt
		}
	}

	return maxNum
}

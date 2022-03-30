package problems

func convert(s string, numRows int) string {
	minRow := 0           //0
	maxRow := numRows - 1 //3
	resIdx := 0
	res := make([]byte, len(s))

	for start := minRow; start <= maxRow; start++ {
		i := start
		firstAdd := 2 * (i - minRow)
		secondAdd := 2 * (maxRow - i)
		if firstAdd == 0 {
			for i < len(s) {
				res[resIdx] = s[i]
				resIdx++
				i += secondAdd
			}
		} else if secondAdd == 0 {
			for i < len(s) {
				res[resIdx] = s[i]
				resIdx++
				i += firstAdd
			}
		} else {
			for i < len(s) {
				// add first
				res[resIdx] = s[i]
				resIdx++
				i += secondAdd
				// add second
				if i >= len(s) {
					break
				}
				res[resIdx] = s[i]
				resIdx++
				i += firstAdd
			}
		}
	}
	return string(res)
}

package problems

func valid(s string) bool {
	count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}
	if count == 0 {
		return true
	}
	return false
}

func generateAll(current string, n int, res *[]string) {
	if len(current) == n {
		// fmt.Println(current)
		if valid(current) {
			*res = append(*res, current)
		}
	} else {
		current += "("
		generateAll(current, n, res)
		current = current[:len(current)-1]
		current += ")"
		generateAll(current, n, res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	generateAll("", n*2, &res)

	return res
}

func back2(cur *string, count int, n int, res *[]string) {
	if len(*cur) == n*2 && count == 0 {
		*res = append(*res, *cur)
	} else {
		*cur += "("
		count++
		if count > n {
			*cur = (*cur)[:len(*cur)-1]
			count--
			return
		}
		back2(cur, count, n, res)
		*cur = (*cur)[:len(*cur)-1]
		count--
		*cur += ")"
		count--
		if count < 0 {
			return
		}
		back2(cur, count, n, res)
		*cur = (*cur)[:len(*cur)-1]
		count++
	}
}

func generateParenthesis2(n int) []string {
	res := []string{}
	cur := ""
	// count := 0
	back2(&cur, 0, n, &res)

	return res
}

func back3(cur *string, open int, close int, n int, res *[]string) {
	if len(*cur) == n*2 {
		*res = append(*res, *cur)
	}

	if open < n {
		*cur += "("
		open++
		back3(cur, open, close, n, res)
		*cur = (*cur)[:len(*cur)-1]
		open--
	}
	if close < open {
		*cur += ")"
		close++
		back3(cur, open, close, n, res)
		*cur = (*cur)[:len(*cur)-1]
		close--
	}
}

func generateParenthesis3(n int) []string {
	res := []string{}
	cur := ""
	// count := 0
	back3(&cur, 0, 0, n, &res)

	return res
}

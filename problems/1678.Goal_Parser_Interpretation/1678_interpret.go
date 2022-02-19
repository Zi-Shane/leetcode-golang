package problems

func interpret(command string) string {
	res := ""
	dict := map[string]string{"G": "G", "()": "o", "(al)": "al"}
	tmp := []rune{}
	for _, ch := range command {
		tmp = append(tmp, ch)
		if v, ok := dict[string(tmp)]; ok {
			res += v
			tmp = []rune{}
		}
	}

	return res
}

func interpret2(command string) string {
	res := ""
	// map[string]string = {"G": "G", "()": "o", "(al)": "al"}
	for i := 0; i < len(command); {
		if command[i] == 'G' {
			res += "G"
			i += 1
		} else if command[i] == '(' && command[i+1] == ')' {
			res += "o"
			i += 2
		} else if command[i] == '(' && command[i+1] == 'a' {
			res += "al"
			i += 4
		}
	}

	return res
}

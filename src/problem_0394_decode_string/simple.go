package solution

func decodeString(s string) string {
	result, _ := decode(s)
	return result
}

func decode(s string) (string, int) {
	result := make([]byte, 0)

	i := 0
	for ; i < len(s); i++ {
		time := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			time = time*10 + int(s[i] - '0')
			i += 1
		}

		if s[i] == '[' {
			sub, next := decode(s[i+1:])
			for j := 0; j < time; j++ {
				result = append(result, sub...)
			}

			i += next
		} else if s[i] == ']' {
			break
		} else {
			result = append(result, s[i])
		}
	}

	return string(result), i+1
}
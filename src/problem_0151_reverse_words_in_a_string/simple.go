package solution

import "strings"

func reverseWords(s string) string {
	var resultBuilder strings.Builder

	end := len(s)
	toContinue := true
	isFirst := true

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			prevEnd := end
			end = i

			if toContinue {
				continue
			}

			if !isFirst {
				resultBuilder.WriteByte(' ')
			} else {
				isFirst = false
			}

			resultBuilder.WriteString(s[i+1 : prevEnd])

			toContinue = true
		} else {
			toContinue = false
		}
	}

	if !toContinue {
		if !isFirst {
			resultBuilder.WriteByte(' ')
		}

		resultBuilder.WriteString(s[:end])
	}

	return resultBuilder.String()
}

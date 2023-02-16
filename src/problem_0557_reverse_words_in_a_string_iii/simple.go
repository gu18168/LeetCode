package solution

import "strings"

// Better: ss := ([]byte)s
func reverseWords(s string) string {
	var builder strings.Builder

	sep := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			continue
		}

		builder.Write(reverseStr(s[sep:i]))
		builder.WriteByte(' ')

		sep = i + 1
	}

	builder.Write(reverseStr(s[sep:]))

	return builder.String()
}

func reverseStr(s string) []byte {
	n := len(s)

	result := make([]byte, n)

	for idx := 0; idx < n/2+1; idx++ {
		result[idx], result[n-1-idx] = s[n-1-idx], s[idx]
	}

	return result
}

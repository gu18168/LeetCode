package solution

import "strconv"

func compress(chars []byte) int {
	result := 0
	start, end := 0, 1

	for ; end < len(chars); end = end + 1 {
		if chars[start] == chars[end] {
			continue
		}

		chars[result] = chars[start]
		result += 1

		if end-start != 1 {
			num := strconv.Itoa(end - start)

			for i := 0; i < len(num); i++ {
				chars[result] = num[i]
				result += 1
			}

		}

		start = end
	}

	chars[result] = chars[start]
	result += 1

	if end-start != 1 {
		num := strconv.Itoa(end - start)

		for i := 0; i < len(num); i++ {
			chars[result] = num[i]
			result += 1
		}

	}

	return result
}

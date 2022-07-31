package solution

import "strconv"

func getHint(secret string, guess string) string {
	bits := [10]int{}

	bulls, cows := 0, 0
	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			bulls += 1
		} else {
            bits[secret[i] - '0'] += 1
        }
	}

	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			continue
		}

		if num := bits[guess[i] - '0']; num > 0 {
			cows += 1
			bits[guess[i] - '0'] -= 1
		}
	}

	results := strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return results
}
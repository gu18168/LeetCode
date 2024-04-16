package solution

func predictPartyVictory(senate string) string {
	numR, numD := 0, 0

	for {
		haveR, haveD := false, false
		queue := make([]byte, 0)

		for i := range senate {
			if senate[i] == 'R' {
				if numD > 0 {
					numD -= 1

					continue
				}

				numR += 1

				if numR > 0 {
					queue = append(queue, 'R')
					haveR = true
				}
			} else if senate[i] == 'D' {
				if numR > 0 {
					numR -= 1

					continue
				}

				numD += 1

				if numD > 0 {
					haveD = true
					queue = append(queue, 'D')
				}
			}
		}

		if haveR && !haveD {
			return "Radiant"
		} else if haveD && !haveR {
			return "Dire"
		}

		numR, numD = -numD, -numR
		senate = string(queue)
	}
}

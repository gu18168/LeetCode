package solution

import "slices"

func reverseWordsInplace(s string) string {
	bs := []byte(s)

	// reverse whole string
	slices.Reverse(bs)

	lastSpacePos, writePos := -1, 0
	toContinue := true
	isFirst := true

	for i := range bs {
		if bs[i] == ' ' {
			spaceStart := lastSpacePos
			lastSpacePos = i

			if toContinue {
				continue
			}

			if isFirst {
				isFirst = false
			} else {
				bs[writePos] = ' '
				writePos += 1
			}

			slices.Reverse(bs[spaceStart+1 : i])

			for j := spaceStart + 1; j < i; j++ {
				bs[writePos] = bs[j]
				writePos += 1
			}

			toContinue = true
		} else {
			toContinue = false
		}
	}

	if !toContinue {
		if !isFirst {
			bs[writePos] = ' '
			writePos += 1
		}

		slices.Reverse(bs[lastSpacePos+1:])

		for j := lastSpacePos + 1; j < len(bs); j++ {
			bs[writePos] = bs[j]
			writePos += 1
		}
	}

	return string(bs[:writePos])
}

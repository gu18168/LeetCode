package solution

func maxArea(height []int) int {
	leftH, rightH := 0, 0
	area := 0

	for start, end := 0, len(height)-1; end > start; {
		if leftH <= rightH {
			for ; end > start; start++ {
				if height[start] > leftH {
					leftH = height[start]

					break
				}
			}
		} else {
			for ; end > start; end-- {
				if height[end] > rightH {
					rightH = height[end]

					break
				}
			}
		}

		if start == end {
			break
		}

		area = max(area, (end-start)*min(leftH, rightH))
	}

	return area
}

package solution

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}

		if flowerbed[i] != 1 {
			n -= 1
		}

		i += 1
	}

	return n <= 0
}

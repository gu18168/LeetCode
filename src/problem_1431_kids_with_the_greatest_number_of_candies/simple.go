package solution

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := slices.Max(candies)

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= maxCandy
	}

	return result
}

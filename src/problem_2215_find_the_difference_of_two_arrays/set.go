package solution

func findDifference(nums1 []int, nums2 []int) [][]int {
	numOfNums1 := make(map[int]struct{})
	numOfNums2 := make(map[int]struct{})

	for _, num := range nums1 {
		numOfNums1[num] = struct{}{}
	}

	for _, num := range nums2 {
		numOfNums2[num] = struct{}{}
	}

	result := make([][]int, 2)

	resultOfNums1 := make(map[int]struct{})
	resultOfNums2 := make(map[int]struct{})

	for _, num := range nums1 {
		if _, ok := numOfNums2[num]; ok {
			continue
		}

		resultOfNums1[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := numOfNums1[num]; ok {
			continue
		}

		resultOfNums2[num] = struct{}{}
	}

	for num := range resultOfNums1 {
		result[0] = append(result[0], num)
	}

	for num := range resultOfNums2 {
		result[1] = append(result[1], num)
	}

	return result
}

package solution

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, num := range nums {
		if ind, ok := indexMap[num]; ok {
			return []int{ind, index}
		}

		indexMap[target - num] = index
	}

	return nil
}
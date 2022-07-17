package solution

func arrayNesting(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	result, res := 0, length
	visited := make(map[int]struct{}, length)

	for i := 0; i < length; i++ {
		if _, ok := visited[i]; ok {
			continue
		}

		now, size := i, 0
		for _, ok := visited[now]; !ok; _, ok = visited[now] {
			visited[now] = struct{}{}
			size++
			now = nums[now]
		}

		if size > result {
			result = size
		}

		res -= size
		if res < result {
			break
		}
	}

	return result
}
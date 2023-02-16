package solution

func moveZerosCount(nums []int) {
	cnt := 0

	for _, num := range nums {
		nums[cnt] = num

		if num != 0 {
			cnt++
		}
	}

	for i := cnt; i < len(nums); i++ {
		nums[i] = 0
	}
}

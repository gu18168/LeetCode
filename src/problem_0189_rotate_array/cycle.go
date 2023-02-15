package solution

func rotateCycle(nums []int, k int) {
	length := len(nums)
	k %= length

	// a 圈，每圈大小 n = an
	// 跳了 b 步（b 个元素），步长为 k = bk
	// 回到起点，所以 an = bk, b = an/k = lcm(n, k) / k
	// n / b = nk / lcm(n, k) = gcd(n, k)
	for start, count := 0, gcd(k, length); start < count; start++ {
		num := nums[start]
		for idx := (start + k) % length; idx != start; idx = (idx + k) % length {
			nums[idx], num = num, nums[idx]
		}

		nums[start] = num
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

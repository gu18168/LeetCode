package solution

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	outdate := -1
	for i := range this.requests {
		if this.requests[i]+3000 >= t {
			break
		}

		outdate = i
	}

	this.requests = this.requests[outdate+1:]

	return len(this.requests)
}

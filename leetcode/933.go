package leetcode

type RecentCounter struct {
	Requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{Requests: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.Requests = append(this.Requests, t)
	for this.Requests[0] < t-3000 {
		this.Requests = this.Requests[1:]
	}
	return len(this.Requests)
}

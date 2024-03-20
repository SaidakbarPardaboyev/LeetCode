type RecentCounter struct {
	ls []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.ls = append(this.ls, t)
	minNum := t - 3000
	for !(this.ls[0] >= minNum) {
		this.ls = this.ls[1:]
	}
	return len(this.ls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
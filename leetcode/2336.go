package leetcode

type SmallestInfiniteSet struct {
	deleted        map[int]bool
	currentMinimal int
}

func Constructor2() SmallestInfiniteSet {
	return SmallestInfiniteSet{deleted: make(map[int]bool), currentMinimal: 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	this.deleted[this.currentMinimal] = true
	x := this.currentMinimal

	for i := this.currentMinimal + 1; ; i++ {
		if _, ok := this.deleted[i]; !ok {
			this.currentMinimal = i
			break
		}
	}

	return x
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	delete(this.deleted, num)
	if num < this.currentMinimal {
		this.currentMinimal = num
	}
}

package main

import "container/heap"

func main() {
	points := [][]int{[]int{1, 2}, []int{3, 4}}
	kClosest(points, 1)
}

type pair struct {
	dist  int
	point []int
}

type hp []pair

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i].dist > h[j].dist }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosest(points [][]int, k int) [][]int {
	h := make(hp, k)
	for i, v := range points[:k] {
		h[i] = pair{v[0]*v[0] + v[1]*v[1], v}
	}

	heap.Init(&h)
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; h[0].dist > dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0)
		}
	}
	var ans [][]int
	for _, p := range h {
		ans = append(ans, p.point)
	}

	return ans

}

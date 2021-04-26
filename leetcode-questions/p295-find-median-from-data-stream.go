/**
 * LeetCode : Find Median from Data Stream
 * https://leetcode.com/problems/find-median-from-data-stream/
 *
 * submission : faster than 65%
 */
package main

import "container/heap"

type MedianFinder struct {
	Small minHeap
	Large minHeap
}
type minHeap []int
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor() MedianFinder {
	m := MedianFinder{}
	heap.Init(&m.Small)
	heap.Init(&m.Large)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.Large, num)
	heap.Push(&this.Small, -heap.Pop(&this.Large).(int))
	if this.Large.Len() < this.Small.Len() {
		heap.Push(&this.Large, -heap.Pop(&this.Small).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Large.Len() > this.Small.Len() {
		return float64(this.Large[0])
	} else {
		return float64(this.Large[0] - this.Small[0]) / 2
	}
}

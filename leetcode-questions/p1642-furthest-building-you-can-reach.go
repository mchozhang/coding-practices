/**
 * LeetCode : Furthest Building You Can Reach
 * https://leetcode.com/problems/furthest-building-you-can-reach/
 *
 * submission : faster than 95%
 */
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 96ms
func furthestBuilding(heights []int, bricks int, ladders int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(minHeap, diff)
		}
		if minHeap.Len() > ladders {
			bricks -= heap.Pop(minHeap).(int)
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 92ms
func furthestBuilding2(heights []int, bricks int, ladders int) int {
	needs := make([]int, len(heights))
	sum := make([]int, len(heights))

	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[i-1] {
			needs[i] = heights[i] - heights[i-1]
		}
		sum[i] = sum[i-1] + needs[i]
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	pos := 0
	for ladders >= 0 {
		start := pos
		high := len(sum) - 1
		for pos < high {
			mid := (pos + high + 1) / 2
			if sum[mid] > bricks {
				high = mid - 1
			} else {
				pos = mid
			}
		}
		if ladders == 0 || pos == len(sum)-1 {
			return pos
		}

		// use a ladder
		pos++
		ladders--

		for i := start + 1; i <= pos; i++ {
			if needs[i] != 0 {
				heap.Push(maxHeap, needs[i])
			}
		}

		bricks += heap.Pop(maxHeap).(int)
	}
	return pos
}

func main() {
	heights := []int{1, 5, 10, 13, 16, 19}
	// 3 4 5
	// ladder: 1 0
	// brick: 8-3, 2
	fmt.Println(furthestBuilding(heights, 8, 2))

}

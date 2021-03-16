/**
 * LeetCode : Jump Game II
 * https://leetcode.com/problems/jump-game-ii/
 *
 * submission : faster than 100%
 */
package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	queue := make([]int, 0)
	queue = append(queue, 0)
	stepQueue := make([]int, 0)
	stepQueue = append(stepQueue, 1)
	var explored  = map[int]int {0:1}
	for len(queue) != 0 {
		size := nums[queue[0]]
		steps := stepQueue[0]
		if queue[0] + size >= len(nums) - 1 {
			return steps
		}
		for i := 1; i <= size; i++ {
			if _, in := explored[queue[0] + i]; !in {
				queue = append(queue, queue[0] + i)
				stepQueue = append(stepQueue, steps + 1)
			}
		}
		queue = queue[1:]
		stepQueue = stepQueue[1:]
	}
	return 0
}



/**
 * 210. Course Schedule II
 * https://leetcode.com/problems/course-schedule-ii/
 *
 * submission: faster than 99%
 */
package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegrees := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		indegrees[pre[0]]++
	}

	order := make([]int, 0)
	queue := []int {}
	for i, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, i)
			order = append(order, i)
		}
	}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range graph[current] {
			indegrees[child]--
			if indegrees[child] == 0 {
				queue = append(queue, child)
				order = append(order, child)
			}
		}
	}
	if len(order) == numCourses {
		return order
	} else {
		return nil
	}
}

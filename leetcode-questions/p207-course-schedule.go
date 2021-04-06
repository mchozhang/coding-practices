/**
 * 10. Course Schedule
 * https://leetcode.com/problems/course-schedule/
 *
 * submission: faster than 100%
 */
package main

/**
 * BFS, 8ms
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	// build graph
	indegrees := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indegrees[prerequisite[0]]++
	}

	var queue []int
	count := 0
	// put 0 degree nodes in queue
	for i, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, i)
			count++
		}
	}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, child := range graph[current] {
			indegrees[child]--
			if indegrees[child] == 0 {
				count++
				queue = append(queue, child)
			}
		}
	}
	return count == numCourses
}

/**
 * non-recursion DFS, 8ms
 */
func canFinish2(numCourses int, prerequisites [][]int) bool {
	// 0: unvisited 1: visiting 2: visited
	visited := make([]int, numCourses)

	// build graph
	graph := buildGraph(numCourses, prerequisites)
	var stack []int

	for i := 0; i < numCourses; i++ {
		if visited[i] == 2 {
			continue
		}

		stack = append(stack, i)
		visited[i] = 1
		for len(stack) != 0 {
			current := stack[len(stack)-1]
			visited[current] = 1
			pop := true
			for _, child := range graph[current] {
				if visited[child] == 1 {
					// cycle found
					return false
				} else if visited[child] == 0 {
					stack = append(stack, child)
					pop = false
				}
			}
			if pop {
				stack = stack[:len(stack)-1]
				visited[current] = 2
			}
		}
	}
	return true
}

func buildGraph(num int, prerequisites [][]int) [][]int {
	graph := make([][]int, num)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

/**
 * build graph and depth first search
 * speed: 4ms
 */
func canFinish3(numCourses int, prerequisites [][]int) bool {
	// 0: unvisited 1: visiting(in stack) 2: visited
	visited := make([]int, numCourses)

	// build graph
	graph := buildGraph(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		if dfs(graph, i, visited) == false {
			return false
		}
	}
	return true
}

/**
 * return whether the node is acyclic, return false if cycle exist
 */
func dfs(graph [][]int, course int, visited []int) bool {
	if visited[course] == 1 {
		return false
	}
	if visited[course] == 2 {
		return true
	}

	visited[course] = 1
	for _, child := range graph[course] {
		if dfs(graph, child, visited) == false {
			return false
		}
	}
	visited[course] = 2
	return true
}

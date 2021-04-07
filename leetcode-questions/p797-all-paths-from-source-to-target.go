/**
 * 796. Rotate String
 * https://leetcode.com/problems/rotate-string/
 *
 * submission: faster than 95%
 */
package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	dfs(&graph, 0, []int{0}, &res)
	return res
}

func dfs(graph [][]int, pos int, path []int, res *[][]int) {
	if pos == len(graph) - 1 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*res = append(*res, newPath)
		return
	}
	for _, child := range graph[pos] {
		dfs(graph, child, append(path, child), res)
	}
}

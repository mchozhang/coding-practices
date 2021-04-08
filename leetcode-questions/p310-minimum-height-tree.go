/**
 * 310. Minimum Height Trees
 * https://leetcode.com/problems/minimum-height-trees/
 *
 * submission: faster than 100%
 */
package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	// put leaves in leaves, iteratively remove leaves
	var leaves []int
	for i := 0; i < n; i++ {
		if len(graph[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for true {
		var nextLayer []int

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				for i, node := range graph[neighbor] {
					if node == leaf {
						graph[neighbor] = append(graph[neighbor][:i], graph[neighbor][i+1:]...)
						if len(graph[neighbor]) == 1 {
							nextLayer = append(nextLayer, neighbor)
						}
						break
					}
				}
			}
		}

		// if all leaves are removed, then current layer is the answer
		if len(nextLayer) == 0 {
			return leaves
		}
		leaves = nextLayer
	}
	return []int{}
}

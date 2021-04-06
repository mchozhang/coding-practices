/**
 * 997. Find the Town Judge
 * https://leetcode.com/problems/find-the-town-judge/
 *
 * submission: faster than 100%
 */
package main

/**
 * uses only 1 array
 */
func findJudge(N int, trust [][]int) int {
	counter := make([]int, N+1)
	for _, edge := range trust {
		counter[edge[1]]++
		counter[edge[0]]--
	}
	for i := 1; i <= N;i++ {
		if counter[i] == N-1 {
			return i
		}
	}
	return -1
}


func findJudge2(N int, trust [][]int) int {
	indegrees := make([]int, N+1)
	outdegree := make([]int, N+1)
	judge := -1
	count := 0
	for _, edge := range trust {
		indegrees[edge[1]]++
		outdegree[edge[0]]++
	}
	for i := 1; i <= N;i++ {
		if indegrees[i] == N-1 && outdegree[i] == 0{
			judge = i
			count++
			if count > 1 {
				return -1
			}
		}
	}
	return judge
}
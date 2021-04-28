/**
 * LeetCode : Keys and Rooms
 * https://leetcode.com/problems/keys-and-rooms/
 *
 * submission : faster than 97%
 */
package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]int{0: 0}

	var stack []int
	stack = append(stack, rooms[0]...)

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[current] = 0
		for _, key := range rooms[current] {
			if _, isIn := visited[key]; !isIn {
				stack = append(stack, key)
			}
		}
	}
	return len(visited) == len(rooms)
}

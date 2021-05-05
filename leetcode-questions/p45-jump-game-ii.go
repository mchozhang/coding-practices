/**
 * LeetCode : Jump Game II
 * https://leetcode.com/problems/jump-game-ii/
 *
 * submission : faster than 100%
 */
package main

func jump(nums []int) int {
    count, curEnd, farthest := 0, 0, 0
    for i := 0; i < len(nums)-1;i++ {
        if nums[i] + i > farthest {
            farthest = nums[i] + i
        }
        if i == curEnd {
            count++
            curEnd = farthest
        }
    }
    return jump
}

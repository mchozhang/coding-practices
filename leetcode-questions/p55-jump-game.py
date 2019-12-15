# LeetCode : Jump Game
# https://leetcode.com/problems/jump-game/
# Given an array of non-negative integers, you are initially positioned at the first index of the array.
#
# Each element in the array represents your maximum jump length at that position.
#
# Determine if you are able to reach the last index.
#
#  Example:
#  Input: [2,3,1,1,4]
#  Output: true
#
#  submission : faster than 99.45%

from typing import List


def canJump(nums: List[int]) -> bool:
    maxReach = 0
    for i, value in enumerate(nums):
        reach = i + value
        if maxReach < i:
            return False

        if reach > maxReach:
            maxReach = reach
    return True


if __name__ == "__main__":
    nums = [1, 1, 2, 2, 0, 1, 1]
    print(canJump(nums))  # true

    nums = [2, 3, 1, 1, 4]
    print(canJump(nums))  # true

    nums = [2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0,
            0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9,
            3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5,
            8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6]
    print(canJump(nums))  # false

    nums = [3, 2, 1, 0, 4]
    print(canJump(nums))  # false

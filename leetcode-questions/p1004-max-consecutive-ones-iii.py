# LeetCode : 219. Contains Duplicate II
# https://leetcode.com/problems/contains-duplicate-ii/

# submission : faster than 99%

from typing import List


def longestOnes(nums: List[int], k: int) -> int:
    window = 0
    zeros = 0

    for i, num in enumerate(nums):
        print(i, num)
        print("window ", window, zeros)
        if num == 0:
            zeros += 1

        if zeros <= k:
            window += 1
        elif nums[i-window] == 0:
            zeros -= 0

    return window


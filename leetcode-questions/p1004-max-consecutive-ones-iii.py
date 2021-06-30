# LeetCode : 1004. Max Consecutive Ones III
# https://leetcode.com/problems/max-consecutive-ones-iii/

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


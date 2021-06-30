# LeetCode : 485. Max Consecutive Ones
# https://leetcode.com/problems/max-consecutive-ones/

# submission : faster than 20%
from typing import List


def findMaxConsecutiveOnes(nums: List[int]) -> int:
    res = 0
    current = 0
    for num in nums:
        if num == 1:
            current += 1
        else:
            current = 0
        res = max(current, res)
    return res

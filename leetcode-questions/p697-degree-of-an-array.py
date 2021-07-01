# LeetCode : 697. Degree of an Array
# https://leetcode.com/problems/degree-of-an-array/

# submission : faster than 71%

from typing import List


def findShortestSubArray(nums: List[int]) -> int:
    maxCount = 0
    res = float('inf')
    counter = dict()
    first = dict()

    for i, num in enumerate(nums):
        if num in counter:
            counter[num] += 1
        else:
            counter[num] = 1
            first[num] = i

        length = i - first[num] + 1
        if counter[num] > maxCount:
            maxCount = counter[num]
            res = length
        elif counter[num] == maxCount:
            res = min(res, length)

    return res


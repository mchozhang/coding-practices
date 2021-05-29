#  1695. Maximum Erasure Value
#  https://leetcode.com/problems/maximum-erasure-value/
#
#  submission : faster than 90%

from typing import List


def maximumUniqueSubarray(nums: List[int]) -> int:
    left = 0
    seen = dict()
    current_sum = 0
    maximum = 0
    for i, num in enumerate(nums):
        current_sum += num
        if num in seen:
            new_left = seen[num] + 1
            for j in range(left, new_left):
                del seen[nums[j]]
                current_sum -= nums[j]
            left = new_left
        seen[num] = i
        maximum = max(maximum, current_sum)
    return maximum

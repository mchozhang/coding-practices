# LeetCode : 219. Contains Duplicate II
# https://leetcode.com/problems/contains-duplicate-ii/

# submission : faster than 20%
from typing import List


def containsNearbyDuplicate(nums: List[int], k: int) -> bool:
    seen = dict()
    for i, num in enumerate(nums):
        if num in seen and i - seen[num] <= k:
            return True
        seen[num] = i
    return False

#  454. 4Sum II
#  https://leetcode.com/problems/4sum-ii/
#
#  submission: faster than 34%
import heapq


def fourSumCount(
    nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]
):
    seen = dict()
    count = 0
    n = len(nums1)
    for num1 in nums1:
        for num2 in nums2:
            key = num1 + num2
            seen[key] = seen.get(key, 0) + 1

    for num1 in nums3:
        for num2 in nums4:
            count += seen.get(-num1 - num2, 0)
    return count
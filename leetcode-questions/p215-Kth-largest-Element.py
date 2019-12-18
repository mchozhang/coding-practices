# LeetCode : Kth Largest Element in an Array
# https://leetcode.com/problems/kth-largest-element-in-an-array/
# Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order,
# not the kth distinct element.
#
# Example:
# Input: [3,2,1,5,6,4] and k = 2
# Output: 5
#
# submission : faster than 99.00%

import heapq
from typing import List


def findKthLargest(nums: List[int], k: int) -> int:
    heapq.heapify(nums)
    return heapq.nlargest(k,nums)[-1]


if __name__ == "__main__":
    nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]
    print(findKthLargest(nums, 4))

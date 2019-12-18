# LeetCode : Sliding Window Maximum
# https://leetcode.com/problems/sliding-window-maximum/
# Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very
# right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return
# the max sliding window.
#
# Example:
# Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
# Output: [3,3,5,5,6,7]
#
# submission : faster than 85.00%

from collections import deque
from typing import List


def maxSlidingWindow(nums: List[int], k: int) -> List[int]:
    if not nums:
        return []

    deq = deque()
    res = []

    for i, num in enumerate(nums):
        while len(deq) != 0:
            if num > deq[-1]:
                deq.pop()
            else:
                break
        deq.append(num)

        if i > k - 1:
            # pop the maximum
            if nums[i - k] == deq[0]:
                deq.popleft()

        res.append(deq[0])

    return res[k - 1:]


if __name__ == "__main__":
    nums = [-7, -8, 7, 8, 3, 1, 6, 0]
    print(maxSlidingWindow(nums, 2))

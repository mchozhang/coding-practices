#  135. Candy
#  https://leetcode.com/problems/candy/
#
#  submission : faster than 99%

from typing import List
import heapq


# O(N)
def candy2(nums: List[int]) -> int:
    up, down, peak = 0, 0, 0
    res = 1
    for i in range(1, len(nums)):
        if nums[i] > nums[i - 1]:
            down = 0
            up += 1
            peak = up
            res += up + 1
        elif nums[i] == nums[i - 1]:
            up = down = peak = 0
            res += 1
        else:
            down += 1
            up = 0
            res += down + (1 if peak < down else 0)
    return res


# O(NlogN)
def candy(nums: List[int]) -> int:
    if len(nums) == 1:
        return 1

    pq = []
    candies = [1 for _ in range(len(nums))]
    heapq.heapify(pq)

    for i, v in enumerate(nums):
        heapq.heappush(pq, (v, i))

    while pq:
        _, pos = heapq.heappop(pq)
        if 0 < pos and nums[pos] > nums[pos - 1]:
            candies[pos] = candies[pos - 1] + 1
        if pos < len(nums) - 1 and nums[pos] > nums[pos + 1] and candies[pos] < candies[pos + 1] + 1:
            candies[pos] = candies[pos + 1] + 1

    return sum(candies)

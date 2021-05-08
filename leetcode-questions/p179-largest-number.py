#   LeetCode : 179. Largest Number
#   https://leetcode.com/problems/largest-number/
#
#   submission : faster than 83%

from typing import List


class Num:
    def __init__(self, value):
        self.s = str(value)

    def __lt__(self, other):
        return other.s + self.s > self.s + other.s

class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        nums = sorted([Num(v) for v in nums], reverse=True)
        if nums[0].s == "0":
            return "0"
        
        res = ""
        for _, num in enumerate(nums):
            res += num.s
        return res


if __name__ == "__main__":
    arr = [1, 10, 122, 11, 121,12,120,2]
    print(largestNumber(arr))

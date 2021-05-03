#  1480. Running Sum of 1d Array
#  https://leetcode.com/problems/running-sum-of-1d-array/
#
#  submission : faster than 96%
import copy


def runningSum(nums: List[int]) -> List[int]:
    for i, num in enumerate(nums[1:]):
        nums[i + 1] += nums[i]
    return nums

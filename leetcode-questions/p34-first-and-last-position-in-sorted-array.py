#  34. Find First and Last Position of Element in Sorted Array
#  https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
#
#  submission : faster than 88.13%

from typing import List


def searchRange(nums: List[int], target: int) -> List[int]:
    if not nums:
        return [-1, -1]

    def binarySearchFirst(low, high):
        while low < high:
            mid = (low + high) // 2
            if nums[mid] >= target:
                high = mid
            else:
                low = mid + 1
        return low if nums[low] == target else -1

    def binarySearchLast(low, high):
        while low < high:
            mid = (low + high + 1) // 2
            if nums[mid] <= target:
                low = mid
            else:
                high = mid - 1

        return low if nums[low] == target else -1

    return [binarySearchFirst(0, len(nums) - 1), binarySearchLast(0, len(nums) - 1)]


if __name__ == "__main__":
    nums = []
    target = 5
    print(searchRange(nums, target))

#  33. Search in Rotated Sorted Array
#  https://leetcode.com/problems/search-in-rotated-sorted-array/
#
#  submission : faster than 6.7%

from typing import List


def search(nums: List[int], target: int) -> int:
    n = len(nums) - 1
    low, high = 0, n
    if n < 0:
        return -1

    # find the position of the smallest
    while low < high:
        mid = (low + high) // 2
        if nums[mid] > nums[high]:
            low = mid + 1
        else:
            high = mid

    # binary search
    if target > nums[n]:
        low, high = 0, low - 1
    else:
        low, high = low, n

    while low < high:
        mid = (low + high) // 2

        if nums[mid] >= target:
            high = mid
        else:
            low = mid + 1

    if nums[low] == target:
        return low
    else:
        return -1


if __name__ == "__main__":
    a = [6,7,8, 0, 1, 2, 3, 4, 5]
    print(search(a, 3))

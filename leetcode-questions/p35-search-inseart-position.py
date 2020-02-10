# 35. Search Insert Position
# https://leetcode.com/problems/search-insert-position/
#
# submission: faster than 92%

from typing import List


def searchInsert(nums: List[int], target: int) -> int:
    if not nums:
        return 0

    if target > nums[-1]:
        return len(nums)

    low, high = 0, len(nums) - 1

    while low < high:
        mid = (low + high) // 2
        if nums[mid] >= target:
            high = mid
        else:
            low = mid + 1

    return low


if __name__ == "__main__":
    nums = [1, 3, 5, 7]
    print(searchInsert(nums, 0))

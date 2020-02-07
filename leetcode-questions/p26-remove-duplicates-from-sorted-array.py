# 26. Remove Duplicates from Sorted Array
# https://leetcode.com/problems/remove-duplicates-from-sorted-array/
#
# submission: faster than 91%

from typing import List


def removeDuplicates(nums: List[int]) -> int:
    if not nums:
        return 0

    count = 1
    i = 1
    n = len(nums)
    while i < n:
        if nums[i] != nums[i - 1]:
            count += 1
            nums[count - 1] = nums[i]
        i += 1
    print(nums)
    return count


if __name__ == "__main__":
    nums = [0, 1, 1, 1, 1, 2, 2, 3, 3, 4]
    print(removeDuplicates(nums))

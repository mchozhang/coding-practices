#  39. First Missing Positive
#  https://leetcode.com/problems/first-missing-positive/
#
#  submission : faster than 78.68%

from typing import List


def firstMissingPositive(nums: List[int]) -> int:
    n = len(nums)
    i = 0
    while i < n:
        value = nums[i]
        if 0 < value <= n and nums[value - 1] != value:
            nums[i], nums[value - 1] = nums[value - 1], nums[i]
            continue
        i += 1

    for i, value in enumerate(nums):
        if value != i + 1:
            return i + 1

    return n + 1


if __name__ == "__main__":
    # nums = [12, 8, 9, 11, 7]
    nums = [1, 1]
    print(firstMissingPositive(nums))

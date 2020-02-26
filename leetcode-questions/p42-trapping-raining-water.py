#  42. Trapping Raining Water
#  https://leetcode.com/problems/trapping-rain-water/
#
#  submission : faster than 85.68%

from typing import List


def trap(height: List[int]) -> int:
    left, right = 0, len(height) - 1
    maxLeft, maxRight = 0, 0
    res = 0

    while left < right:
        if height[left] < height[right]:
            if height[left] > maxLeft:
                maxLeft = height[left]
            res += maxLeft - height[left]
            left += 1
        else:
            if height[right] > maxRight:
                maxRight = height[right]
            res += maxRight - height[right]
            right -= 1

    return res


if __name__ == "__main__":
    nums = [3, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 1]
    print(trap(nums))

    nums = [5,1]
    print(trap(nums))

    nums = [5, 1, 1, 2, 3, 1, 2, 1, 5]
    print(trap(nums))

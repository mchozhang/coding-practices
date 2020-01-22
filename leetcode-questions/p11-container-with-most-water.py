#  11. Container With Most Water
#  https://leetcode.com/problems/container-with-most-water/
#
#  submission : faster than 63%

from typing import List


def maxArea(height: List[int]) -> int:
    left, right = 0, len(height) - 1
    area = min(height[left], height[right]) * (right - left)
    while left < right:
        if height[left] < height[right]:
            left += 1
        elif height[left] == height[right]:
            left += 1
            right -= 1
        else:
            right -= 1

        newArea = min(height[left], height[right]) * (right - left)
        area = max(newArea, area)

    return area


if __name__ == "__main__":
    a = [1, 8, 6, 2, 5, 4, 8, 3, 7]

    print(maxArea(a))

#  1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
#  https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
#
#  submission : faster than 5%
from typing import List


def maxArea(h: int, w: int, horizontalCuts: List[int], verticalCuts: List[int]) -> int:
    horizontalCuts.sort()
    verticalCuts.sort()
    height = max(horizontalCuts[0], h - horizontalCuts[-1])
    width = max(verticalCuts[0], w - verticalCuts[-1])
    for i in range(1, len(horizontalCuts)):
        height = max(height, horizontalCuts[i] - horizontalCuts[i - 1])

    for i in range(1, len(verticalCuts)):
        width = max(width, verticalCuts[i] - verticalCuts[i - 1])
    return (width * height) % 1000000007

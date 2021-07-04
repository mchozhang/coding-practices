# LeetCode : 658. Find K Closest Elements
# https://leetcode.com/problems/find-k-closest-elements/

# submission : faster than 68%

from typing import List
import bisect


# binary search, 296ms
def findClosestElements2(arr: List[int], k: int, x: int) -> List[int]:
    closest = bisect.bisect_right(arr, x, 0, len(arr) - 1)
    if closest > 0 and arr[closest] - x >= x - arr[closest - 1]:
        closest -= 1

    left, right = closest - 1, closest + 1
    for i in range(k - 1):
        if right == len(arr):
            left -= 1
        elif left == -1:
            right += 1
        elif x - arr[left] <= arr[right] - x:
            left -= 1
        else:
            right += 1
    return arr[left + 1:right]

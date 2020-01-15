#  4. Median of Two Sorted Arrays
#  https://leetcode.com/problems/median-of-two-sorted-arrays/
#  There are two sorted arrays nums1 and nums2 of size m and n respectively.
#  Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
#  You may assume nums1 and nums2 cannot be both empty.
#
#  Example:
#  input: nums1 = [1, 2] nums2 = [3, 4]
#  Output: (2 + 3)/2 = 2.5
#
#  submission : faster than 20%

from typing import List


def findMedianSortedArrays(nums1: List[int], nums2: List[int]) -> float:
    """
    A median divides the array into 2 equal length parts, where the right part is always greater than the left.
    Ensures:
    1. len(left) == len(right) (or len(right) + 1)
    2. max(left) <= min(right)
    then the median is (max(left) + min(right)) / 2
    Search i, j so that A[i] >= B[j-1] and A[i-1] <= B[j], j = (m + n + 1)/2 - i
    """
    # m is the shorter list length, n is the longer list length
    arr1, m = nums1, len(nums1)
    arr2, n = nums2, len(nums2)

    if m > n:
        arr1, arr2 = arr2, arr1
        m, n = n, m

    low, high = 0, m
    half_len = (m + n + 1) // 2
    while low <= high:
        i = (low + high) // 2
        j = half_len - i

        if i < m and arr2[j - 1] > arr1[i]:
            low = i + 1
        elif i > 0 and arr1[i - 1] > arr2[j]:
            high = i - 1
        else:
            if i == 0:
                max_left = arr2[j - 1]
            elif j == 0:
                max_left = arr1[i - 1]
            else:
                max_left = max(arr1[i - 1], arr2[j - 1])

            if (m + n) % 2 == 1:
                return max_left

            if i == m:
                min_right = arr2[j]
            elif j == n:
                min_right = arr1[i]
            else:
                min_right = min(arr1[i], arr2[j])

            return (max_left + min_right) / 2


if __name__ == "__main__":
    a = [1, 2, 3, 4, 6, 7, 8]
    b = [3, 5]

    print(findMedianSortedArrays(a, b))

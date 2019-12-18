# LeetCode : Merge Intervals
# https://leetcode.com/problems/merge-intervals/
# Given a collection of intervals, merge all overlapping intervals.
#
# Example:
# Input: [[1,3],[2,6],[8,10],[15,18]]
# Output: [[1,6],[8,10],[15,18]]
#
# submission : faster than 98.59%

from typing import List


def merge(intervals: List[List[int]]) -> List[List[int]]:
    if len(intervals) < 2:
        return intervals
    intervals.sort(key=lambda x: x[0])

    res = [intervals[0]]
    for interval in intervals[1:]:
        last = res[-1]
        start, end = last[0], last[1]

        if interval[0] <= end < interval[1]:
            # overlap
            last[1] = interval[1]
        elif interval[0] > end:
            # apart
            res.append(interval)

    return res


if __name__ == "__main__":
    nums = [[2, 5], [1, 3], [1, 2], [2, 6]]
    print(merge([[1, 4], [2, 5]]))
    print(merge(nums))

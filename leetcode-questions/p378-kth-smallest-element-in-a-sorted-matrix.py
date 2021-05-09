#  378. Kth Smallest Element in a Sorted Matrix
#  https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
#
#  submission : faster than 43%

import heapq


def kthSmallest(matrix: List[List[int]], k: int) -> int:
    h = []
    columnIndexs = []
    heapq.heapify(h)
    for i, v in enumerate(matrix[0]):
        columnIndexs.append(0)
        h.append((v, i))
    
    res = matrix[0][0]
    for i in range(k):
        res, col = heapq.heappop(h)
        columnIndexs[col] += 1
        if columnIndexs[col] < len(matrix):
            heapq.heappush(h, (matrix[columnIndexs[col]][col], col))

    return res

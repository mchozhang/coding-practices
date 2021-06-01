#  1480. Running Sum of 1d Array
#  https://leetcode.com/problems/running-sum-of-1d-array/
#
#  submission : faster than 91%

from typing import List

from itertools import product


def maxAreaOfIsland(grid: List[List[int]]) -> int:
    res = 0
    n, m = len(grid), len(grid[0])
    dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
    for i, j in product(range(n), range(m)):
        if grid[i][j] == 1:
            queue = [(i, j)]
            grid[i][j] = 2
            area = 1
            while queue:
                x, y = queue.pop(0)
                for dir in dirs:
                    nextX, nextY = x + dir[0], y + dir[1]
                    if -1 < nextX < n and -1 < nextY < m and \
                            grid[nextX][nextY] == 1:
                        area += 1
                        grid[nextX][nextY] = 2
                        queue.append((nextX, nextY))

            res = max(res, area)
    return res

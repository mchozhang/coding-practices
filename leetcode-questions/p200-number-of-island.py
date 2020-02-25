#  200. Number of Island
#  https://leetcode.com/problems/number-of-islands/
#
#  submission : faster than 88.55%

from typing import List


def numIslands(grid: List[List[str]]) -> int:
    if not grid:
        return 0

    res = 0
    rowNum = len(grid)
    colNum = len(grid[0])

    def search(row, col):
        if -1 < row < rowNum and -1 < col < colNum and grid[row][col] == '1':
            grid[row][col] = '#'
            search(row + 1, col)
            search(row - 1, col)
            search(row, col + 1)
            search(row, col - 1)

    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if cell == '1':
                search(i, j)

                res += 1
    print(grid)
    return res


if __name__ == "__main__":
    grid = [
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "1"]
    ]
    print(numIslands(grid))

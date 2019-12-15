#  LeetCode : Spiral Matrix
#  https://leetcode.com/problems/spiral-matrix/
#  Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
#
#  Example:
#  Input: [
#  [ 1, 2, 3 ],
#  [ 4, 5, 6 ],
#  [ 7, 8, 9 ]
# ]
#  Output: [1,2,3,6,9,8,7,4,5]
#
#  submission : faster than 95.82%
from typing import List


def spiralOrder(matrix: List[List[int]]) -> List[int]:
    if len(matrix) == 0 or len(matrix[0]) == 0:
        return []

    res = []
    x = y = 0
    m = len(matrix)  # length of column
    n = len(matrix[0])  # length of row
    visited = [[0 for i in range(n)] for j in range(m)]
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    currentDir = 0  # index of directions

    def isValid(x, y):
        return 0 <= x < m and 0 <= y < n and visited[x][y] == 0

    while True:
        res.append(matrix[x][y])
        visited[x][y] = 1
        nextX = x + directions[currentDir][0]
        nextY = y + directions[currentDir][1]

        if isValid(nextX, nextY):
            x = nextX
            y = nextY
        else:
            # take a turn
            currentDir += 1
            currentDir %= 4
            nextX = x + directions[currentDir][0]
            nextY = y + directions[currentDir][1]
            if isValid(nextX, nextY):
                x = nextX
                y = nextY
            else:
                # finish
                break
    return res


if __name__ == "__main__":
    matrix = [
        [1, 2, 3, 1],
        [4, 5, 6, 2],
        [7, 8, 9, 3]
    ]
    print(spiralOrder(matrix))

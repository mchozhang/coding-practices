#  289. Game of Life
#  https://leetcode.com/problems/game-of-life/
#
#  submission : faster than 41%


def gameOfLife(board: List[List[int]]) -> None:
    dirs = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]
    for i in range(len(board)):
        for j in range(len(board[0])):
            live, dead = 0, 0
            for dir in dirs:
                nextX, nextY = i + dir[0], j + dir[1]
                if -1 < nextX < len(board) and -1 < nextY < len(board[0]):
                    if board[nextX][nextY] == 1 or board[nextX][nextY] == 2:
                        live += 1
                    elif board[nextX][nextY] == 0 or board[nextX][nextY] == 3:
                        dead += 1
            if board[i][j] == 0 and live == 3:
                board[i][j] = 3
            elif board[i][j] == 1 and (live < 2 or live > 3):
                board[i][j] = 2

    for i in range(len(board)):
        for j in range(len(board[0])):
            if board[i][j] == 2:
                board[i][j] = 0
            elif board[i][j] == 3:
                board[i][j] = 1

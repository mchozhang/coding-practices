#  79. Word Search
#  https://leetcode.com/problems/word-search/
#
#  submission : faster than 90%

from typing import List


def exist(board: List[List[str]], word: str) -> bool:
    """
    DFS
    """
    if not str or not board:
        return False

    wordLen = len(word) - 1
    rowNum = len(board)
    colNum = len(board[0])

    def search(row, col, step) -> bool:
        if -1 < row < rowNum and -1 < col < colNum and board[row][col] == word[step]:
            # current cell is valid
            if step == wordLen:
                return True

            tmp = board[row][col]
            board[row][col] = '\0'
            if search(row + 1, col, step + 1) or \
                    search(row - 1, col, step + 1) or \
                    search(row, col + 1, step + 1) or \
                    search(row, col - 1, step + 1):
                return True
            board[row][col] = tmp

        return False

    for i, row in enumerate(board):
        for j, c in enumerate(row):
            if board[i][j] == word[0] and search(i, j, 0):
                return True

    return False


if __name__ == "__main__":
    board = [
        ['A', 'B', 'C', 'E'],
        ['S', 'F', 'C', 'S'],
        ['A', 'D', 'E', 'E']
    ]
    board = [
        ["A", "B", "C", "E"],
        ["S", "F", "E", "S"],
        ["A", "D", "E", "E"]
    ]
    # board = [['a', 'a', 'a']]

    word = "ABCCED"
    word = "ABCESEEEFS"
    # word = "aaa"
    print(exist(board, word))

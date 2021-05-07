/**
 * LeetCode : Surrounded Regions
 * https://leetcode.com/problems/surrounded-regions/
 * 
 * submission: faster than 40%
 */

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solve = function (board) {
    let m = board.length, n = board[0].length
    dirs = [[1, 0], [0, 1], [-1, 0], [0, -1]]

    let checkEdge = (i, j) => {
        let queue = [i * n + j]
        board[i][j] = 'o'
        while (queue.length !== 0) {
            let pos = queue.shift()
            let x = Math.floor(pos / n), y = pos % n
            dirs.forEach(([dx, dy]) => {
                let nextX = x + dx, nextY = y + dy
                if (-1 < nextX && nextX < m && -1 < nextY && nextY < n) {
                    if (board[nextX][nextY] === 'O') {
                        board[nextX][nextY] = 'o'
                        queue.push(nextX * n + nextY)
                    }
                }
            })
        }
    }


    for (let i = 0; i < n; i++) {
        if (board[0][i] === 'O') {
            checkEdge(0, i)
        }
        if (board[m - 1][i] === 'O') {
            checkEdge(m - 1, i)
        }
    }
    for (let i = 1; i < m - 1; i++) {
        if (board[i][0] === 'O') {
            checkEdge(i, 0)
        }
        if (board[i][n - 1] === 'O') {
            checkEdge(i, n - 1)
        }
    }

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (board[i][j] === 'O') {
                board[i][j] = 'X'
            } else if (board[i][j] === 'o') {
                board[i][j] = 'O'
            }
        }
    }
};

let grid = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
solve(grid)
console.log(grid)
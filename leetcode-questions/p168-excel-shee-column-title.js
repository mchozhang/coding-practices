/**
 * LeetCode : 168. Excel Sheet Column Title
 * https://leetcode.com/problems/excel-sheet-column-title/
 *
 * submission : faster than 68%
 */

var convertToTitle = function (num) {
    let title = ""
    let codeA = 'A'.charCodeAt(0)
    for (; num != 0; num = Math.floor((num - 1) / 26)) {
        title = String.fromCharCode((num - 1) % 26 + codeA) + title
    }
    return title
};
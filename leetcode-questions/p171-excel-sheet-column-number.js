/**
 * LeetCode : Excel Sheet Column Number
 * https://leetcode.com/problems/excel-sheet-column-number/
 *
 * submission : faster than 86%
 */

var titleToNumber = function (columnTitle) {
    let base = 0
    let codeA = 'A'.charCodeAt(0) - 1
    for (let i = 0; i < columnTitle.length; i++) {
        base = base * 26 + columnTitle.charCodeAt(i) - codeA
    }
    return base
};
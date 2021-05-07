/**
 * LeetCode : 387. First Unique Character in a String
 * https://leetcode.com/problems/first-unique-character-in-a-string/
 *
 * submission : faster than 99%
 */

/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function (s) {
    let counter = new Array(26).fill(0)
    for (let i = 0; i < s.length; i++) {
        counter[s.charCodeAt(i) - 'a'.charCodeAt(0)]++
    }
    for (let i = 0; i < s.length; i++) {
        if (counter[s.charCodeAt(i) - 'a'.charCodeAt(0)] == 1) {
            return i
        }
    }
    return -1
};
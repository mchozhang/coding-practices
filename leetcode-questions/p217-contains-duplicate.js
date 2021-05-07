/**
 * LeetCode : Contains Duplicate
 * https://leetcode.com/problems/contains-duplicate/
 *
 * submission : faster than 70%
 */

var containsDuplicate = function (nums) {
    let seen = new Set()
    for (let num of nums) {
        if (seen.has(num)) {
            return true
        }
        seen.add(num)
    }
    return false
};
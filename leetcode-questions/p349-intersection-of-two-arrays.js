/**
 * LeetCode : 349. Intersection of Two Arrays
 * https://leetcode.com/problems/intersection-of-two-arrays/
 *
 * submission : faster than 76%
 */

var intersection = function(nums1, nums2) {
    let intersection = new Set(nums1)
    let res = []
    nums2.forEach(num => {
        if (intersection.has(num)) {
            res.push(num)
            intersection.delete(num)
        }
    });
    return res
};
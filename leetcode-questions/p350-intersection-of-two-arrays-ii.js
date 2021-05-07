/**
 * LeetCode : 350. Intersection of Two Arrays II
 * https://leetcode.com/problems/intersection-of-two-arrays-ii/
 *
 * submission : faster than 82%
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function (nums1, nums2) {
    let dict = {}
    let res = []
    nums1.forEach(num => {
        if (num in dict) {
            dict[num]++
        } else {
            dict[num] = 1
        }
    });

    nums2.forEach(num => {
        if (dict[num] > 0) {
            dict[num]--
            res.push(num)
        }
    });
    return res
};
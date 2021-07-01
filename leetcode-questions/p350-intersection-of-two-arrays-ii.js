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


/**
 * 2 sorted arrays
 * @param nums1
 * @param nums2
 * @returns {*[]}
 */
var intersect2 = function (nums1, nums2) {
    let i = 0, j = 0
    let res = []
    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] === nums2[j]) {
            res.push(nums1[i])
            i++
            j++
        } else if (nums1[i] < nums2[j]) {
            i++
        } else {
            j++
        }
    }

    return res
};
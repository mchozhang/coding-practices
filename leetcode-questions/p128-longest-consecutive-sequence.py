# 128. Longest Consecutive Sequence
# https://leetcode.com/problems/longest-consecutive-sequence/
# Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
# Your algorithm should run in O(n) complexity.
#
# Example:
# Input: [100, 4, 200, 1, 3, 2]
# Output: 4
# Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
#
# submission: faster than 90.43%


def longestConsecutive(nums) -> int:
    # seen[num] represents the length of the longest consecutive sequence containing the value num
    seen = dict()
    longest = 0

    for num in nums:
        if num in seen:
            continue

        right = left = 0
        if num + 1 in seen:
            right = seen[num + 1]
        if num - 1 in seen:
            left = seen[num - 1]

        sum = 1 + right + left
        seen[num] = sum

        if num + right in seen:
            seen[num + right] = sum
        if num - left in seen:
            seen[num - left] = sum

        if sum > longest:
            longest = sum

    return longest


if __name__ == "__main__":
    nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
    print(longestConsecutive(nums))

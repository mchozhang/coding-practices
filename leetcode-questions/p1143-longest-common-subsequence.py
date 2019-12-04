# 1143. Longest Common Subsequence
# https://leetcode.com/problems/longest-common-subsequence/
# Given two strings text1 and text2, return the length of their longest common subsequence.
# A subsequence of a string is a new string generated from the original string with some characters(can be none)
# deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while
# "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
# If there is no common subsequence, return 0.
# Example:
# Input: text1 = "abcde", text2 = "ace"
# Output: 3
# Explanation: The longest common subsequence is "ace" and its length is 3.


class Solution:

    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        """
        time complexity: O(MN)
        """
        dp = dict()

        def get_dp(i, j):
            if (i, j) in dp:
                return dp[(i, j)]
            else:
                return 0

        for i in range(1, len(text1) + 1):
            for j in range(1, len(text2) + 1):
                if text1[i - 1] == text2[j - 1]:
                    dp[(i, j)] = 1 + get_dp(i - 1, j - 1)
                else:
                    dp[(i, j)] = max(get_dp(i - 1, j), get_dp(i, j - 1))

        return dp[len(text1), len(text2)]


if __name__ == "__main__":
    s1 = "abc"
    s2 = "aabdc"

    print(Solution().longestCommonSubsequence(s1, s2))

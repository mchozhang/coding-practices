# 5. Longest Palindromic Substring
# https://leetcode.com/problems/longest-palindromic-substring/
# Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
#
# Example:
# Input: "babad"
# Output: "bab"
#
# submission: faster than 90.85%


def longestPalindrome(s: str) -> str:
    if s is None or s == "":
        return ""

    longestLength = 0
    res = ""
    for i, c in enumerate(s):

        if isPalindrome(s, i - 1 - longestLength, i):
            res = s[i - 1 - longestLength: i + 1]
            longestLength += 2
        elif isPalindrome(s, i - longestLength, i):
            res = s[i - longestLength: i + 1]
            longestLength += 1

    return res


def isPalindrome(s, begin, end):
    if begin < 0:
        return False

    low, high = begin, end
    while low < high:
        if s[low] != s[high]:
            return False
        low += 1
        high -= 1
    return True


if __name__ == "__main__":
    s = "ababab"
    print(longestPalindrome(s))

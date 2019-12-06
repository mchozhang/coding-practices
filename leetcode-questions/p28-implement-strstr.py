# 28. Implement strStr()
# https://leetcode.com/problems/implement-strstr/
# Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
#
# Example:
# Input: haystack = "hello", needle = "ll"
# Output: 2
#
# submission: faster than 92%
# todo: KMP algorithm


def strStr(haystack: str, needle: str) -> int:
    """
    time complexity: O(MN)
    """

    if not needle or len(needle) == 0:
        return 0

    for i, c in enumerate(haystack):
        if i + len(needle) > len(haystack):
            break
        if haystack[i:i + len(needle)] == needle:
            return i
    return -1


if __name__ == "__main__":
    print(strStr("abccd", "b"))

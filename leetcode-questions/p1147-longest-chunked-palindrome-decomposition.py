#  1147. Longest Chunked Palindrome Decomposition
#  https://leetcode.com/problems/longest-chunked-palindrome-decomposition/
#
#  submission : faster than 75%

def longestDecomposition(text: str) -> int:
    low, high = 0, len(text)-1
    count = 0
    s1, s2 = "", ""
    while low < high:
        s1 += text[low]
        s2 = text[high] + s2

        if s1 == s2:
            s1, s2 = "", ""
            count += 2
        low += 1
        high -= 1
    if low == high or s1 != s2:
        count += 1

    return count

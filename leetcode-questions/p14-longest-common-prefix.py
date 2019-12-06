# 14. Longest Common Prefix
# https://leetcode.com/problems/longest-common-prefix/
# Write a function to find the longest common prefix string amongst an array of strings.
# If there is no common prefix, return an empty string "".
# Example:
# Input: ["flower","flow","flight"]
# Output: "fl"
#
# submission: faster than 89%


def longestCommonPrefix(strs) -> str:
    if len(strs) == 0:
        return ""

    prefix = ""
    for i, c in enumerate(strs[0]):
        is_common = True
        for str in strs[1:]:
            if i >= len(str) or str[i] != c:
                is_common = False
                break

        if is_common:
            prefix += c
        else:
            break
    return prefix


if __name__ == "__main__":
    # strs = ["flower", "flow", "flight"]
    strs = ["lower", "low", "lojer"]
    print(longestCommonPrefix(strs))


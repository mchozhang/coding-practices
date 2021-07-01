# LeetCode : 278. First Bad Version
# https://leetcode.com/problems/first-bad-version/

# submission : faster than 80%

# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
def isBadVersion(version):
    return True

# 1 1 0
def firstBadVersion(n):
    low = 1
    high = n
    while low < high:
        mid = (low + high) // 2
        if isBadVersion(mid):
            high = mid
        else:
            low = mid + 1

    return low

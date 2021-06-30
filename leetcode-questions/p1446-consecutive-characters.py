# 1446. Consecutive Characters
# https://leetcode.com/problems/consecutive-characters/

# submission : faster than 64%

def maxPower(s: str) -> int:
    res = 1
    prev = s[0]
    current = 1
    for letter in s[1:]:
        if letter == prev:
            current += 1
        else:
            current = 1
        prev = letter
        res = max(res, current)
    return res

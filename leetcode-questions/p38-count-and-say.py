# 38. Count and Say
# https://leetcode.com/problems/count-and-say/
#
# submission: faster than 97%

def countAndSay(n: int) -> str:
    s = "1"
    for i in range(2, n + 1):
        newStr = ""
        count = 1
        pre = s[0]
        for c in s[1:]:
            if c == pre:
                count += 1
            else:
                newStr += chr(ord('0') + count) + pre
                count = 1
            pre = c
        newStr += chr(ord('0') + count) + pre
        s = newStr
    return s

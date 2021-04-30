#  91. Decode Ways
#  https://leetcode.com/problems/decode-ways/
#
#  submission : faster than 100%

def numDecodings(s: str) -> int:
    if s[0] == '0':
        return 0
    dp = [1]
    for i, letter in enumerate(s):
        value = dp[i]
        if letter == '0':
            if "00" < s[i - 1:i + 1] < "30":
                value = dp[i - 1]
            else:
                return 0

        elif i > 0 and "10" < s[i - 1:i + 1] < "27":
            value += dp[i - 1]
        dp.append(value)
    return dp[-1]

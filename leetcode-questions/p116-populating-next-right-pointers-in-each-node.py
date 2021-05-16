#  65. Valid Number
#  https://leetcode.com/problems/valid-number/
#
#  submission : faster than 53%

def isNumber(s: str) -> bool:
    s = s.lower()
    nums = s.split('e')
    if len(nums) == 1:
        return isDecimal(nums[0]) or isInteger(nums[0])
    if len(nums) == 2:
        return (isDecimal(nums[0]) or isInteger(nums[0])) and isInteger(nums[1])
    return False


def isDecimal(s):
    if not s or '.' not in s:
        return False

    nums = s.split('.')
    if nums[0] and (nums[0][0] == "+" or nums[0][0] == "-"):
        nums[0] = nums[0][1:]

    if len(nums) != 2 or (not nums[0] and not nums[1]):
        return False

    return isDigitStr(nums[0]) and isDigitStr(nums[1])


def isInteger(s):
    if s and (s[0] == "+" or s[0] == "-"):
        s = s[1:]
    if not s:
        return False
    return isDigitStr(s)


def isDigitStr(s):
    i = 0
    while i < len(s):
        if not '0' <= s[i] <= '9':
            return False
        i += 1
    return True

#  43. Multiply Strings
#  https://leetcode.com/problems/multiply-strings/
#
#  submission: faster than 35%


def multiply(num1: str, num2: str) -> str:
    if num1 == "0" or num2 == "0":
        return "0"

    nums = [0] * (len(num1) + len(num2))

    for i, c1 in enumerate(num1[::-1]):
        last = 0
        for j, c2 in enumerate(num2[::-1]):
            last = (
                last // 10
                + (ord(c1) - ord("0")) * (ord(c2) - ord("0"))
                + nums[-1 - i - j]
            )
            nums[-1 - i - j] = last % 10
        if last > 9:
            nums[-1 - i - len(num2)] = last // 10

    return "".join(map(str, nums)).lstrip("0")

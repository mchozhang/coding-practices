# 28. Pow(x, n)
# https://leetcode.com/problems/powx-n/
# Implement pow(x, n), which calculates x raised to the power n (xn).
#
# Example:
# Input: 2.00, 10
# Output: 1024.00
#
# submission: faster than 64.80%


def myPow(x: float, n: int) -> float:
    if n < 0:
        x = 1 / x
        n = -n

    def pow(x, n):
        if n == 0:
            return 1
        return pow(x * x, n // 2) if n % 2 == 0 else x * pow(x * x, n // 2)

    return pow(x, n)


if __name__ == "__main__":
    print(myPow(2.1, 3))


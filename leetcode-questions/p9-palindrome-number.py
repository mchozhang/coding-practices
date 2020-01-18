# 9. Palindrome Number
# https://leetcode.com/problems/palindrome-number/
# Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
#
# Example:
# Input: 121
# Output: true
#
# submission: faster than 93%


def isPalindrome(x: int) -> bool:
    if x < 0:
        return False

    s = str(x)
    halfLen = len(s) // 2
    i = 0
    while i < halfLen:
        if s[i] != s[- i - 1]:
            return False
        i += 1
    return True


def isPalindrome2(x: int) -> bool:
    """
    implementation without converting into string
    """
    if x < 0:
        return False

    reverse, original = 0, x

    while x != 0:
        reverse = reverse * 10 + x % 10
        x //= 10

    return reverse == original


if __name__ == "__main__":
    print(isPalindrome2(101))


#  17. Letter Combinations of a Phone Number
#  https://leetcode.com/problems/letter-combinations-of-a-phone-number/
#
#  submission : faster than 86%

from typing import List


def letterCombinations(digits: str) -> List[str]:
    if digits == "":
        return []

    keyboard = {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    res = []

    def combination(step, selected):
        if step == len(digits) - 1:
            for c in keyboard[digits[step]]:
                res.append(selected + c)
        else:
            for c in keyboard[digits[step]]:
                combination(step + 1, selected + c)

    combination(0, "")
    return res


if __name__ == "__main__":
    print(letterCombinations("423"))

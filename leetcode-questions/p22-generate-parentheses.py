# 22. Generate Parentheses
# https://leetcode.com/problems/generate-parentheses/
#
# submission: faster than 90%

from typing import List


def generateParenthesis(n: int) -> List[str]:
    res = []

    def parentheses(left, right, selected, stack):
        if left == 0 and right > 0:
            res.append(selected + ')' * right)
        else:
            if stack:
                parentheses(left, right - 1, selected + ")", stack[:-1])

            parentheses(left - 1, right, selected + "(", stack + ['('])

    parentheses(n, n, "", [])
    return res


if __name__ == "__main__":
    print(generateParenthesis(3))

#  LeetCode : Longest Valid Parentheses
#  https://leetcode.com/problems/3sum/
#  Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed)
#  parentheses substring.
#
#  Example:
#  Input: "(()"
#  Output: 2
#
#  submission : faster than 99.44%


def longestValidParentheses(s: str) -> int:
    stack = []  # store the index of the '('
    validPair = dict()

    for i, c in enumerate(s):
        if c == '(':
            stack.append(i)
        elif stack:
            # a valid match occurs
            matchIndex = stack.pop()
            # check if the new match results in a longer valid match
            if matchIndex - 1 in validPair:
                validPair[i] = validPair[matchIndex - 1] + (i - matchIndex + 1)
            else:
                validPair[i] = i - matchIndex + 1

    return max(validPair.values()) if validPair else 0


if __name__ == "__main__":
    print(longestValidParentheses("(())()(()(("))
    print(longestValidParentheses("(()(((()"))
    print(longestValidParentheses("((("))
    print(longestValidParentheses("((((()()(())"))
    print(longestValidParentheses(")()())"))
    print(longestValidParentheses(")()())(()))((((()))))())"))

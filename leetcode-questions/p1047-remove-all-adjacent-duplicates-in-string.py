# LeetCode : 219. Contains Duplicate II
# https://leetcode.com/problems/contains-duplicate-ii/

# submission : faster than 65%

def removeDuplicates(s: str) -> str:
    stack = []
    for letter in s:
        if len(stack) > 0 and stack[-1] == letter:
            stack.pop()
        else:
            stack.append(letter)
    return "".join(stack)
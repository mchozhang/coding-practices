#  227. Basic Calculator II
#  https://leetcode.com/problems/basic-calculator-ii/
#
#  submission : faster than 97%


# one stack
def calculate(s: str) -> int:
    s = s.replace(" ", "")
    s += "+"
    stack = []
    num = 0
    sign = "+"
    for c in s:
        if '0' <= c <= '9':
            num = num * 10 + ord(c) - ord('0')
        else:
            if sign == '+':
                stack.append(num)
            elif sign == '-':
                stack.append(-num)
            elif sign == "*":
                stack.append(num * stack.pop())
            elif sign == '/':
                stack.append(int(float(stack.pop()) / num))
            sign = c
            num = 0

    return sum(stack)


# two stacks
def calculate2(s: str) -> int:
    s = s.replace(" ", "")
    numStack, operatorStack = [], []
    i = 0
    isNum = True
    while i < len(s):
        if isNum:
            num = ""
            while i < len(s) and '0' <= s[i] <= '9':
                num += s[i]
                i += 1

            num = int(num)
            if len(operatorStack) != 0 and \
                    (operatorStack[-1] == '*' or operatorStack[-1] == '/'):
                op = operatorStack.pop()
                num2 = numStack.pop()
                if op == '*':
                    numStack.append(num2 * num)
                else:
                    numStack.append(num2 // num)
            else:
                numStack.append(num)

        else:
            if s[i] != " ":
                operatorStack.append(s[i])

            i += 1
        isNum = not isNum

    res = 0
    while len(numStack) != 0:
        if len(operatorStack) != 0 and operatorStack.pop() == '-':
            res -= numStack.pop()
        else:
            res += numStack.pop()
    return res

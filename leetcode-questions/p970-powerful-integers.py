#  970. Powerful Integers
#  https://leetcode.com/problems/powerful-integers/
#
#  submission : faster than 97%

def powerfulIntegers(x: int, y: int, bound: int) -> List[int]:
    if x == 1 and y == 1:
        if bound > 1:
            return [2]
        else:
            return []
    elif x == 1 or y == 1:
        big = y if x == 1 else x
        res = []
        base = 1
        while base + 1 <= bound:
            res.append(base + 1)
            base *= big
        return res

    res = set()
    left, right = 1, 1
    while left < bound:
        while left + right <= bound:
            res.add(left + right)
            right *= y
        left, right = left * x, 1
    return list(res)

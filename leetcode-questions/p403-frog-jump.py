# 403. Frog Jump
# https://leetcode.com/problems/frog-jump/
# Given a list of stones' positions (in units) in sorted ascending order, determine if the frog is able to cross the
# river by landing on the last stone. Initially, the frog is on the first stone and assume the first jump must be 1
# unit.
#
# Example:
# Input: [0,1,3,5,6,8,12,17]
# Output: true
#
# submission: faster than 68.75%

from typing import List


def canCross(stones: List[int]) -> bool:
    stack = [(0, 0)]
    visited = set()
    while stack:
        curPos, jump = stack.pop()
        curStones = stones[curPos]
        visited.add((curPos, jump))
        if curPos >= len(stones) - 1:
            return True

        i = curPos + 1
        while i < len(stones):
            nextStone = stones[i]
            if curStones + jump - 1 <= nextStone <= curStones + jump + 1 and (i, nextStone - curStones) not in visited:
                stack.append((i, nextStone - curStones))
            elif curStones + jump - 1 > nextStone:
                i += 1
                continue
            else:
                break
            i += 1

    return False


if __name__ == "__main__":
    nums = [0, 1, 3, 5, 6, 8, 12, 17]
    print(canCross(nums))

    nums = [0, 1, 3, 6, 10, 15, 16, 21]
    print(canCross(nums))

    nums = [0, 1, 3, 6, 7]
    print(canCross(nums))

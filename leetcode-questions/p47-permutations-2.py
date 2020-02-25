#  47. Permutations II
#  https://leetcode.com/problems/permutations-ii
#  submission : faster than 64.19%

from typing import List


def permute(nums: List[int]) -> List[List[int]]:
    n = len(nums)
    if n < 2:
        return [nums]

    nums.sort()
    used = []
    for num in nums:
        used.append(False)

    res = []
    selected = []

    def permutation():
        if len(selected) == n:
            res.append(selected.copy())
            return

        for i, value in enumerate(nums):
            if used[i] or (i > 0 and value == nums[i - 1] and used[i - 1] == False):
                continue

            selected.append(value)
            used[i] = True
            permutation()
            selected.pop()
            used[i] = False

    permutation()
    return res


if __name__ == "__main__":
    nums = [-1, -1, 0, 0, 1, 1, 2]
    res = permute(nums)
    print(res)

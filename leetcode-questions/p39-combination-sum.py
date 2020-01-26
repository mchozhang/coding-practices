#  39. Combination Sum
#  https://leetcode.com/problems/combination-sum
#
#  submission : faster than 92.86%

from typing import List


def combinationSum(candidates: List[int], target: int) -> List[List[int]]:
    res = []
    candidates.sort()

    def combination(start, end, tar, selected):
        for index in range(start, end):
            value = candidates[index]
            if value == tar:
                res.append(selected + [value])
                return

            if value > tar:
                return

            combination(index, end, tar - value, selected + [value])

    combination(0, len(candidates), target, [])
    return res


if __name__ == "__main__":
    nums = [2,3,5]
    target = 8
    print(combinationSum(nums, target))

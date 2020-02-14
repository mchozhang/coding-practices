#  40. Combination Sum II
#  https://leetcode.com/problems/combination-sum-ii/
#
#  submission : faster than 93.86%

from typing import List


def combinationSum2(candidates: List[int], target: int) -> List[List[int]]:
    res = []
    candidates.sort()

    def combination(start, tar, selected):
        for index, value in enumerate(candidates[start:]):
            index += start

            if index > start and value == candidates[index - 1]:
                continue

            if value == tar:
                res.append(selected + [value])
                return

            if value > tar:
                return

            combination(index + 1, tar - value, selected + [value])

    combination(0, target, [])
    return res


if __name__ == "__main__":
    nums = [10, 1, 2, 7, 6, 1, 5]
    target = 8
    print(combinationSum2(nums, target))

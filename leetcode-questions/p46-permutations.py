#  46. Permutations
#  https://leetcode.com/problems/permutations
#  submission : faster than 95%

from typing import List


def permute(nums: List[int]) -> List[List[int]]:
    n = len(nums)
    if n < 2:
        return [nums]

    res = []

    def permutation(begin):
        if begin == n - 1:
            res.append(nums.copy())
            return

        for i, value in enumerate(nums[begin:]):
            i += begin
            nums[i], nums[begin] = nums[begin], nums[i]
            permutation(begin + 1)
            nums[i], nums[begin] = nums[begin], nums[i]

    permutation(0)
    return res


if __name__ == "__main__":
    nums = [1,2,3]
    print(permute(nums))

#  18. 4Sum
#  https://leetcode.com/problems/4sum/
#
#  submission : faster than 97%

from typing import List

count = 0


def fourSum(nums: List[int], target: int) -> List[List[int]]:
    nums.sort()
    res = []

    def findNSum(begin, end, n, tar, selected):
        if n < 2 or end - begin + 1 < n or tar < nums[begin] * n or tar > nums[end] * n:
            return

        s = sum(selected) + nums[begin]
        if 0 < target < s or target <= 0 < s:
            return

        elif n == 2:
            while begin < end:
                twoSum = nums[begin] + nums[end]
                if twoSum == tar:
                    res.append(selected + [nums[begin], nums[end]])
                    begin += 1
                    end -= 1
                    while begin < end and nums[begin] == nums[begin - 1]:
                        begin += 1
                    while begin < end and nums[end] == nums[end + 1]:
                        end -= 1
                elif twoSum < tar:
                    begin += 1
                else:
                    end -= 1
        else:
            for i in range(begin, end + 1):
                value = nums[i]
                if i > begin and value == nums[i - 1]:
                    continue

                findNSum(i + 1, end, n - 1, tar - value, selected + [value])

    findNSum(0, len(nums) - 1,  4, target, [])

    return res


if __name__ == "__main__":
    # nums = [-5, -4, -3, -2, 1, 3, 3, 5]
    nums = [-1, 0, -5, -2, -2, -4, 0, 1, -2]

    a = -9
    print(fourSum(nums, a))

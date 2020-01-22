#  LeetCode : 3Sum
#  https://leetcode.com/problems/3sum/
#  Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
#  Find all unique triplets in the array which gives the sum of zero.
#
#  Example:
#  Input: [-1, 0, 1, 2, -1, -4]
#  Output: [ [-1, 0, 1], [-1, -1, 2]]
#
#  submission : faster than 63%


def threeSum(nums):
    res = []
    nums.sort()

    for i, v in enumerate(nums[:-2]):
        if v > 0:
            break

        if i > 0 and v == nums[i - 1]:
            continue
        low = i + 1
        high = len(nums) - 1

        while low < high:
            target = 0 - v
            if nums[low] + nums[high] == target:
                res += [[v, nums[low], nums[high]]]
                low += 1
                high -= 1
                while low < high and nums[low] == nums[low - 1]:
                    low += 1
                while low < high and nums[high] == nums[high + 1]:
                    high -= 1
            elif nums[low] + nums[high] < target:
                low += 1
            else:
                high -= 1
    return res


if __name__ == "__main__":
    a = [-1, 0, 1, 2, -1, -4]

    print(threeSum(a))

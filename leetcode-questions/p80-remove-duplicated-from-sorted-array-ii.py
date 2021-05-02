#  80. Remove Duplicates from Sorted Array II
#  https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
#
#  submission : faster than 76%

def removeDuplicates(nums: List[int]) -> int:
    if len(nums) < 3:
        return len(nums)
    i = 2
    for num in enumerate(nums[2:]):
        if num > nums[i - 2]:
            nums[i] = num
            i += 1
    return i

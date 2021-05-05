#  665. Non-decreasing Array
#  https://leetcode.com/problems/non-decreasing-array/
#
#  submission : faster than 41%

def checkPossibility(nums: List[int]) -> bool:
    changed = False
    for i in range(1, len(nums)):
        if nums[i] < nums[i-1]:
            if changed:
                return False
            changed = True
            if i < len(nums)-1:
                if nums[i-1] > nums[i+1]:
                    # nums[i-1] needs to change
                    nums[i-1] = nums[i]
                    if i > 1 and nums[i-1] < nums[i-2]:
                        return False
                else:
                    # nums[i] needs to change
                    nums[i] = nums[i-1]
    return True


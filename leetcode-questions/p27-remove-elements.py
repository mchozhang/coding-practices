# 27. Remove Element
# https://leetcode.com/problems/remove-element/
#
# submission: faster than 84%

from typing import List


def removeElement(nums: List[int], val: int) -> int:
    tail = len(nums)
    i = 0
    while i < tail:
        if nums[i] == val:
            tail -= 1
            nums[i], nums[tail] = nums[tail], nums[i]
        else:
            i += 1

    return tail

if __name__ == "__main__":
    nums = [3, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 1]
    nums = [0,1,2,2,3,0,4,2]
    print(removeElement(nums, 2))

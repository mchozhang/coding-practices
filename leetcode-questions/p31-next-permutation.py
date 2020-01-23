#  31. Next Permutation
#  https://leetcode.com/problems/next-permutation/
#
#  submission : faster than 67%

from typing import List


def nextPermutation(nums: List[int]) -> None:
    n = len(nums)
    i = n - 1
    reverseStart = 0

    while i > 0:
        if nums[i] > nums[i - 1]:
            j = n - 1

            while j > i - 1:
                if nums[j] > nums[i - 1]:
                    nums[j], nums[i - 1] = nums[i - 1], nums[j]
                    reverseStart = i
                    break
                j -= 1
            break
        i -= 1

    # reverse sort, swap head and tail
    k = reverseStart
    while k < (reverseStart + n) // 2:
        nums[k], nums[reverseStart + n - k - 1] = nums[reverseStart + n - k - 1], nums[k]
        k += 1


if __name__ == "__main__":
    a = [1,2,5,3,1]

    nextPermutation(a)
    print(a)

#  1207. Unique Number of Occurrences
#  https://leetcode.com/problems/unique-number-of-occurrences/
#
#  submission : faster than 100%

def uniqueOccurrences(self, arr: List[int]) -> bool:
    counter = dict()
    for num in arr:
        counter[num] = counter.get(num, 0) + 1

    occurrences = set()
    occurrences.update(counter.values())
    return len(occurrences) == len(counter)

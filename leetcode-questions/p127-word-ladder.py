#  127. Word Ladder
#  https://leetcode.com/problems/word-ladder/
#
#  submission : faster than 5%


def ladderLength(beginWord: str, endWord: str, wordList: List[str]) -> int:
    wordSet = set(wordList)
    queue = set({word for word in wordSet if isAdjacent(beginWord, word)})
    for word in queue:
        wordSet.remove(word)
    length = 1
    while queue:
        nextQueue = set()
        length += 1
        for word in queue:
            if word == endWord:
                return length
            neighbors = {neighbor for neighbor in wordSet if isAdjacent(word, neighbor)}
            nextQueue.update(neighbors)
            for neighbor in neighbors:
                wordSet.remove(neighbor)

        queue = nextQueue
    return 0


def isAdjacent(word1: str, word2: str) -> bool:
    diff = False
    i = 0
    while i < len(word1):
        if word1[i] != word2[i]:
            if diff:
                return False
            diff = True
        i += 1
    return diff

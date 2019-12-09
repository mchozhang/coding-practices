# 269. Alien Dictionary
# https://leetcode.com/problems/alien-dictionary/
# There is a new alien language which uses the latin alphabet. However, the order among letters are unknown to you.
# You receive a list of non-empty words from the dictionary, where words are sorted lexicographically by the rules
# of this new language. Derive the order of letters in this language.
#
# Example:
# Input: ["wrt", "wrf", "er", "ett", "rftt"]
# Output: "wertf"
from collections import Counter
from queue import Queue


def alien_order(words) -> str:
    # build letter set
    letters = set()
    for word in words:
        for letter in word:
            letters.add(letter)

    # build the graph that indicates the order relationship between letters
    # every letter is a node and every pair is an edge of a graph
    graph = set()
    for i, word in enumerate(words[:-1]):
        next_word = words[i + 1]
        min_length = min(len(word), len(next_word))
        j = 0
        while j < min_length:
            if word[j] != next_word[j]:
                graph.add((word[j], next_word[j]))
                break
            j += 1

        # invalid dictionary
        if j == min_length and j < len(word):
            return ""

    # compute the depth of each node through the edges
    incoming = Counter()
    for edge in graph:
        incoming[edge[1]] += 1

    queue = Queue()
    result = ""

    # nodes with no incoming edge can be seen as start point
    for letter in letters:
        if letter not in incoming:
            queue.put(letter)
            result += letter

    # BFS
    while not queue.empty():
        letter = queue.get()
        for edge in graph:
            if letter == edge[0]:
                incoming[edge[1]] -= 1

                # find the new start node
                if incoming[edge[1]] == 0:
                    queue.put(edge[1])
                    result += edge[1]

    # letters not in the graph can be added in any sequence
    for letter in letters:
        if letter not in result:
            result += letter

    return result


if __name__ == "__main__":
    words = ["wrt", "wrf", "er", "ett", "rftt"]
    print(alien_order(words))  # wertf

    words = ["aa", "ab", "ae", "ee", "ec", "ed", "dc"]
    print(alien_order(words))  # abecd

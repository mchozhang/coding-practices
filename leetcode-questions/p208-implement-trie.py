#  2018. Implement Trie
#  https://leetcode.com/problems/implement-trie-prefix-tree/
#
#  submission : faster than 50%

from collections import defaultdict


class TrieNode:
    def __init__(self):
        self.isWord = False
        self.children = defaultdict(TrieNode)


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        current = self.root
        for letter in word:
            current = current.children[letter]
        current.isWord = True

    def search(self, word: str) -> bool:
        current = self.root
        for letter in word:
            current = current.children.get(letter)
            if current is None:
                return False
        return current.isWord

    def startsWith(self, prefix: str) -> bool:
        current = self.root
        for letter in prefix:
            current = current.children.get(letter)
            if current is None:
                return False
        return True


if __name__ == "__main__":
    trie = Trie()
    trie.insert("abc")
    trie.insert("abd")
    trie.insert("hello")

    print(trie.search("abc"))  # true
    print(trie.search("abd"))  # true
    print(trie.search("ab"))  # false
    print(trie.search("abe")) # false

    print(trie.startsWith("a"))  # true
    print(trie.startsWith("hello"))  # true

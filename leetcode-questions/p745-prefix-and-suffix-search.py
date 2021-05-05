#  745. Prefix and Suffix Search
#  https://leetcode.com/problems/prefix-and-suffix-search/
#
#  submission : faster than 5%

from typing import List

class WordFilter:
    def __init__(self, words: List[str]):
        self.words = words

    def f(self, prefix: str, suffix: str) -> int:

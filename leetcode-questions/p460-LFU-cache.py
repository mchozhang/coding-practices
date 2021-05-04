#  460. LFU Cache
#  https://leetcode.com/problems/lfu-cache/
#
#  submission : faster than 99%

from collections import defaultdict, OrderedDict


class Node:
    def __init__(self, value):
        self.value = value
        self.count = 1


class LFUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.table = dict()
        self.counter = defaultdict(lambda: OrderedDict())
        self.minFrequency = 1

    def get(self, key: int) -> int:
        if key not in self.table:
            return -1

        # put node into the next frequency ordered dict
        node = self.table[key]
        if node.count == self.minFrequency and len(self.counter[node.count]) == 1:
            self.minFrequency += 1
        node.count += 1
        self.counter[node.count - 1].pop(key)
        self.counter[node.count][key] = node
        return node.value

    def put(self, key: int, value: int) -> None:
        if self.cap == 0:
            return

        if self.get(key) != -1:
            self.table[key].value = value
            return

        # remove the head node of the min frequency list
        if len(self.table) == self.cap:
            k, head = self.counter[self.minFrequency].popitem(last=False)
            self.table.pop(k)
            self.minFrequency = 1

        node = Node(value)
        self.table[key] = node
        self.counter[1][key] = node
        self.minFrequency = 1

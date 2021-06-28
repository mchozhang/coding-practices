#!/usr/bin/env python
# -*- coding: utf-8 -*-

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []
        self.minValue = None

    def push(self, val: int) -> None:
        if len(self.stack) == 0:
            self.minValue = val
            self.stack.append(0)
        else:
            self.stack.append(val - self.minValue)
            self.minValue = min(val, self.minValue)

    def pop(self) -> None:
        value = self.stack.pop()
        if value < 0:
            self.minValue -= value

    def top(self) -> int:
        value = self.stack[-1]
        if value < 0:
            return self.minValue
        else:
            return self.minValue + value

    def getMin(self) -> int:
        return self.minValue

# implement a queue using 2 stacks


class Queue:
    """
    a queue implementation by using 2 stacks.
    In python a list can be seen as a stack.
    """

    def __init__(self):
        # storage stack
        self.storage = []

        # temporary stack
        self.temp = []

    def push(self, value):
        """
        append an element to the queue bottom
        """
        self.storage.append(value)

    def pop(self):
        """
        get and pop the first element of the queue
        """
        if len(self.storage) == 0:
            return None

        while len(self.storage) > 1:
            self.temp.append(self.storage.pop())

        result = self.storage.pop()

        while len(self.temp) > 0:
            self.storage.append(self.temp.pop())

        return result


if __name__ == "__main__":
    queue = Queue()
    queue.push(1)
    queue.push(2)
    queue.push(3)
    queue.push(4)

    print(queue.pop())
    print(queue.pop())
    print(queue.pop())
    print(queue.pop())
    print(queue.pop())

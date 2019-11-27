# reverse linked list


class Node:
    def __init__(self, value, next_node=None):
        self.value = value
        self.next_node = next_node


def reverse_linked_list(head):
    """
    reverse the linked list
    :param head: head of the linked list
    :return: head node of the reversed linked list
    """
    current = head
    prev = None
    while current is not None:
        next = current.next_node
        current.next_node = prev
        prev = current
        current = next
    head = prev
    return head


if __name__ == "__main__":
    # initialize linked list 1->2->3->4
    head = Node(1, None)
    head.next_node = Node(2, None)
    head.next_node.next_node = Node(3, None)
    head.next_node.next_node.next_node = Node(4, None)

    head = reverse_linked_list(head)

    # output 4 3 2 1
    while head is not None:
        print(head.value)
        head = head.next_node

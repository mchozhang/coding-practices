# AVL tree implementation

from queue import Queue


class Node:
    """
    node in binary search tree
    """

    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

        # the maximum depth that this node has extended, the minimum value is 1
        self.height = 1

        # balance factor: right child height - left child height
        self.balance_factor = 0

    def update(self):
        """
        update the height and balance factor
        """
        left = right = 0
        if self.right:
            right = self.right.height

        if self.left:
            left = self.left.height

        self.height = 1 + max(left, right)

        # compute the balance factor of the node
        self.balance_factor = right - left

class AVLTree:
    """
    An AVL tree is a binary search tree which has the following properties:
    1. The sub-trees of every node differ in height by at most one.
    2. Every sub-tree is an AVL tree.
    This maintains an O(logN) search time. Addition and deletion operations also take O(logN) time.
    """

    def __init__(self):
        self.root = None

    def insert(self, value):
        """
        insert a node to an AVL tree:
        begin the recursive insert to find the insert position
        :param value: value of the new node
        """
        self.root = self.recursive_insert(value, self.root)

    def recursive_insert(self, value, node):
        """
        insert a node to an AVL tree:
        1. Perform the normal BST insertion
        2. Re-balance the tree by performing appropriate rotations on the subtree
        :param value: value of the new node
        :param node: root node of the sub-tree that performs inserting
        :return : the new root node
        """
        # traverse the tree to find the inserting point
        if node is None:
            return Node(value)
        elif value <= node.value:
            node.left = self.recursive_insert(value, node.left)
        else:
            node.right = self.recursive_insert(value, node.right)

        # update the height and balance factor of the visited nodes
        node.update()

        # check whether the node is unbalance, that is balance factor isn't in [-1, 1]
        # left left case: y is left child of z and x is left child of y
        #         z                                      y
        #        / \                                   /   \
        #       y   T4      Right Rotate (z)          x      z
        #      / \          - - - - - - - - ->      /  \    /  \
        #     x   T3                               T1  T2  T3  T4
        #    / \
        #  T1   T2
        if node.balance_factor < -1 and value <= node.left.value:
            return self.right_rotate(node)

        # left right case: y is left child of z and x is right child of y
        #         z                             z                                x
        #        / \                           / \                             /   \
        #       y   T4     Left Rotate (y)    x   T4     Right Rotate (z)     y     z
        #      / \       - - - - - - - - ->  / \        - - - - - - - - ->   / \   /  \
        #     T1   x                        y  T3                           T1 T2 T3  T4
        #         / \                      / \
        #       T2   T3                   T1 T2
        if node.balance_factor < -1 and value > node.left.value:
            node.left = self.left_rotate(node.left)
            return self.right_rotate(node)

        # right right case: y is right child of z and x is right child of y
        #         z                                      y
        #        / \                                   /   \
        #       T1  y       left Rotate (z)           z     x
        #          / \     - - - - - - - - ->        / \   / \
        #         T2  x                             T1 T2 T3  T4
        #            / \
        #          T3   T4
        if node.balance_factor > 1 and value > node.right.value:
            return self.left_rotate(node)

        # right left case: y is right child of z and x is left child of y
        #         z                               z                               x
        #        / \                             / \                            /   \
        #       T1   y     Right Rotate (y)    T1   x     Right Rotate (z)     z     y
        #           / \   - - - - - - - - ->       / \   - - - - - - - - ->   / \   /  \
        #          x   T4                         T2  y                     T1  T2 T3  T4
        #         / \                                / \
        #       T2   T3                             T3 T4
        if node.balance_factor > 1 and value <= node.right.value:
            node.right = self.right_rotate(node.right)
            return self.right_rotate(node)

        # tree is balance after inserting
        return node

    @staticmethod
    def left_rotate(z):
        """
        left rotate the unbalanced node
        :param z: the node performing rotation
        :return: the node after perform left rotation
        """
        y = z.right
        T2 = y.left

        # perform rotation
        y.left = z
        z.right = T2

        # update height
        z.update()
        y.update()

        return y

    @staticmethod
    def right_rotate(z):
        """
        right rotate the unbalanced node
        :param z: the node performing rotation
        :return: the node after perform right rotation
        """
        y = z.left
        T3 = y.right

        # perform rotation
        y.right = z
        z.left = T3

        # update height
        z.update()
        y.update()

        return y


    def level_traverse(self):
        """
        top-down traverse the tree level by level
        """
        if not self.root:
            print("tree is null")
            return []

        queue = Queue()
        queue.put(self.root)
        traverse_sequence = []

        while not queue.empty():
            current = queue.get()
            traverse_sequence.append(current.value)
            if current.left:
                queue.put(current.left)
            if current.right:
                queue.put(current.right)

        return traverse_sequence

    def print_tree(self):
        """
        print the tree
        """
        queue = Queue()
        queue.put(self.root)
        level_dict = dict()
        while not queue.empty():
            current = queue.get()
            current_level_list = level_dict.get(current.level)
            if current_level_list:
                current_level_list.append(current.value)
            else:
                current_level_list = [current.value]

            level_dict[current.level] = current_level_list

            if current.left:
                queue.put(current.left)
            if current.right:
                queue.put(current.right)

        i = 0
        while level_dict.get(i) is not None:
            print(level_dict.get(i))
            i += 1

    @staticmethod
    def get_height(node):
        """
        get height of node
        :return: height
        """
        if node:
            return node.get_height
        else:
            return 0


if __name__ == "__main__":
    tree = AVLTree()
    tree.insert(5)
    tree.insert(4)
    tree.insert(6)
    tree.insert(8)
    tree.insert(1)
    tree.insert(9)
    print(tree.level_traverse())

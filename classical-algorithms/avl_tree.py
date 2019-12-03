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

    def insert(self, key):
        """
        insert a node to an AVL tree:
        1. Perform the normal BST insertion
        2. Re-balance the tree by performing appropriate rotations on the subtree
        :param key: value of the new node
        """
        def recursive_insert(key, node):
            # traverse the tree to find the inserting point
            if node is None:
                return Node(key)
            elif key <= node.value:
                node.left = recursive_insert(key, node.left)
            else:
                node.right = recursive_insert(key, node.right)

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
            if node.balance_factor < -1 and key <= node.left.value:
                return self.right_rotate(node)

            # left right case: y is left child of z and x is right child of y
            #         z                             z                                x
            #        / \                           / \                             /   \
            #       y   T4     Left Rotate (y)    x   T4     Right Rotate (z)     y     z
            #      / \       - - - - - - - - ->  / \        - - - - - - - - ->   / \   /  \
            #     T1   x                        y  T3                           T1 T2 T3  T4
            #         / \                      / \
            #       T2   T3                   T1 T2
            if node.balance_factor < -1 and key > node.left.value:
                node.left = self.left_rotate(node.left)
                return self.right_rotate(node)

            # right right case: y is right child of z and x is right child of y
            #         z                                      y
            #        / \                                   /   \
            #       T1  y       Left Rotate (z)           z     x
            #          / \     - - - - - - - - ->        / \   / \
            #         T2  x                             T1 T2 T3  T4
            #            / \
            #          T3   T4
            if node.balance_factor > 1 and key > node.right.value:
                return self.left_rotate(node)

            # right left case: y is right child of z and x is left child of y
            #         z                               z                               x
            #        / \                             / \                            /   \
            #       T1   y     Right Rotate (y)    T1   x     Left Rotate (z)     z     y
            #           / \   - - - - - - - - ->       / \   - - - - - - - - ->   / \   /  \
            #          x   T4                         T2  y                     T1  T2 T3  T4
            #         / \                                / \
            #       T2   T3                             T3 T4
            if node.balance_factor > 1 and key <= node.right.value:
                node.right = self.right_rotate(node.right)
                return self.left_rotate(node)

            # tree is balance after inserting
            return node

        self.root = recursive_insert(key, self.root)

    def delete(self, key):
        """
        delete the first node containing the value
        :param key: value of the key
        """
        def recursive_delete(key, node):
            if not node:
                return None

            elif key == node.value:
                # find the key
                if not node.left:
                    return node.right

                elif not node.right:
                    return node.left

                # left child will replace the vacancy
                temp = self.get_max_node(node.left)
                node.value = temp.value
                node.left = recursive_delete(temp.value, node.left)

            elif key < node.value:
                node.left = recursive_delete(key, node.left)

            else:
                node.right = recursive_delete(key, node.right)

            if not node:
                return None

            # update height and balance factor
            node.update()

            # left left case
            if node.balance_factor < -1 and node.left.balance_factor <= 0:
                return self.right_rotate(node)

            # left right case
            if node.balance_factor < -1 and node.left.balance_factor > 0:
                node.left = self.left(node.left)
                return self.right_rotate(node)

            # right right case
            if node.balance_factor > 1 and node.right.balance_factor >= 0:
                return self.left_rotate(node)

            # right left case
            if node.balance_factor > 1 and node.right.balance_factor < 0:
                node.right = self.right_rotate(node.right)
                return self.left_rotate(node)

            return node

        self.root = recursive_delete(key, self.root)

    def search(self, value):
        """
        search a value in the tree
        :param value: key value
        :return: first node containing the value
        """
        def recursive_search(node, value):
            if value == node.value:
                return node
            elif node.left and value < node.value:
                return recursive_search(node.left, value)
            elif node.right and value > node.value:
                return recursive_search(node.right, value)

            return None

        return recursive_search(self.root, value)

    def get_min_node(self, node):
        """
        find the node with the minimum value,
        pick the node with the minimum height if their values are equal
        :param node: root node of the sub-tree
        :return: node with the minimum value in the tree
        """
        if node is None or node.left is None:
            return node
        else:
            return self.get_min_node(node.left)

    def get_max_node(self, node):
        """
        find the node with the maximum value from the sub-tree
        :param node: root node of the sub-tree
        :return: node with the maximum value
        """
        if node is None or node.right is None:
            return node
        else:
            return self.get_max_node(node.right)

    @staticmethod
    def left_rotate(a):
        """
        left rotate the unbalanced node
             A                  B
              \                /  \
               B    ------->  A    D
              / \              \
             C   D              C
        :param a: the node A performing left rotation
        :return: the node B after perform left rotation
        """
        b = a.right
        c = b.left

        # perform rotation
        b.left = a
        a.right = c

        # update height
        a.update()
        b.update()

        return b

    @staticmethod
    def right_rotate(a):
        """
        right rotate the unbalanced node
             A                 B
            /                /  \
           B    ------->    C   A
          / \                  /
         C   D                D
        :param a: the node A performing rotation
        :return: the node node_d after perform right rotation
        """
        b = a.left
        d = b.right

        # perform rotation
        b.right = a
        a.left = d

        # update height
        a.update()
        b.update()

        return b

    def pre_order_traverse(self):
        """
        pre-order traversal:
        root node -> left sub-tree -> right sub-tree
        :return: list of node
        """
        nodes = []

        def recursive_preorder(node):
            if node:
                nodes.append(node.value)
                recursive_preorder(node.left)
                recursive_preorder(node.right)

        recursive_preorder(self.root)
        return nodes

    def in_order_traverse(self):
        """
        in-order traversal:
        left sub-tree -> root node -> right sub-tree
        :return: list of node
        """
        nodes = []

        def recursive_in_order(node):
            if node:
                recursive_in_order(node.left)
                nodes.append(node.value)
                recursive_in_order(node.right)

        recursive_in_order(self.root)
        return nodes

    def post_order_traverse(self):
        """
        post-order traversal:
        left sub-tree ->  right sub-tree -> root node
        :return: list of node
        """
        nodes = []

        def recursive_post_order(node):
            if node:
                recursive_post_order(node.left)
                recursive_post_order(node.right)
                nodes.append(node.value)

        recursive_post_order(self.root)
        return nodes

    def level_traverse(self):
        """
        top-down traverse the tree level by level
        """
        if not self.root:
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

    def __repr__(self):
        """
        present the tree with list of (index, value) pair,
        where index represents the node index in the tree
        """
        if not self.root:
            return str([])

        nodes = dict()
        nodes[self.root] = 0

        queue = Queue()
        queue.put(self.root)
        while not queue.empty():
            current = queue.get()
            current_index = nodes[current]

            if current.left:
                queue.put(current.left)
                nodes[current.left] = current_index * 2 + 1

            if current.right:
                queue.put(current.right)
                nodes[current.right] = current_index * 2 + 2

        sorted_nodes = sorted(nodes.items(), key=lambda kv: kv[1])
        node_list = [(index, node.value) for node, index in sorted_nodes]
        return str(node_list)


if __name__ == "__main__":
    tree = AVLTree()
    tree.insert(3)
    tree.insert(8)
    tree.insert(4)
    tree.insert(5)
    tree.insert(6)
    tree.insert(1)
    tree.insert(9)
    print(tree)
    print(tree.pre_order_traverse())
    print(tree.in_order_traverse())
    print(tree.post_order_traverse())
    print(tree.level_traverse())

    # search
    node = tree.search(4)
    if node:
        print(node.value, node.height)
    else:
        print("value doesn't exist")

    # delete
    tree.delete(10)
    tree.delete(0)
    print(tree)
    tree.delete(5)
    tree.delete(8)
    tree.delete(4)
    tree.delete(9)
    tree.delete(6)
    tree.delete(1)
    print(tree)

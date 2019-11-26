# heap sort
# 1. build a heap, max heap for ascending order, min heap for descending order
# 2. swap the root node with the last leaf node one by one


def heapify(arr, len, i):
    """
    heapify the sub-tree root at index i, so that every node will be larger than its children node
    :param arr: array
    :param len: length of arr that needs to be heapified
    :param i: current index
    """
    left = 2 * i + 1
    right = 2 * i + 2

    max_index = i

    # compare current node with its left child
    if len > left and arr[left] > arr[i]:
        max_index = left

    # compare the max with right child
    if len > right and arr[right] > arr[max_index]:
        max_index = right

    if max_index != i:
        # swap the max value to the current node
        arr[i], arr[max_index] = arr[max_index], arr[i]

        # heapifying has to continue in the sub-tree if a swap has been made
        heapify(arr, len, max_index)


def heap_sort(arr):
    """
    ascending order heap sort
    :param arr: array
    """
    # build max heap,
    # start with the last non-leaf node, from right to left and bottom-up
    i = int(len(arr) / 2) - 1
    while i >= 0:
        heapify(arr, len(arr), i)
        i -= 1

    # extract the max value remained the heap one by one
    i = len(arr) - 1
    while i > 0:
        # swap the max value to the bottom
        arr[0], arr[i] = arr[i], arr[0]

        # heapify the rest of the heap
        heapify(arr, i, 0)
        i -= 1


def nth_largest_value(arr, n):
    """
    find the nth largest value in the array:
    :param arr: array
    :param n: 0 < n <= len(arr)
    :return: nth largest value
    """
    # build max heap
    start = int(len(arr) / 2) - 1
    for i in range(start, -1, -1):
        heapify(arr, len(arr), i)

    # extract the max value remained the heap one by one,
    # repeat n times
    start = len(arr) - 1
    for i in range(start, start - n, -1):
        # swap the max value to the bottom
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)

    return arr[-n]


if __name__ == "__main__":
    arr = [6, 5, 7, 0, 1, 8, 2, 9, 3, 4]
    heap_sort(arr)
    print(arr)

    # find 2nd largest
    arr = [6, 5, 7, 0, 1, 8, 2, 9, 3, 4]
    second_largest = nth_largest_value(arr, 2)
    print(second_largest)

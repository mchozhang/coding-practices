# find missing items in the second array
# for example:
# a1 = [1,2,3,4,5]
# a2 = [1,2,3]
# output: [4,5]


def find_missing_items(a1, a2):
    """
    use set to find a list of missing item,
    time complexity: O(m + n)
    space complexity: O(m + n)
    :param a1: first array
    :param a2: second array
    :return: list of missing items
    """
    diff_set = set(a1) - set(a2)

    return list(diff_set)


def find_the_only_missing_item(a1, a2):
    """
    use exclusive or or find the only missing item,
    pre-condition: only 1 item is miss, no duplicated item in array 1
    note: n ^ n = 0, n ^ 0 = n
    time complexity: O(m + n)
    space complexity: O(1)
    :return: the only missing item
    """
    xor = 0
    for num in a1:
        xor ^= num

    for num in a2:
        xor ^= num
    return xor


if __name__ == "__main__":
    a1 = [4, 6, 8, 8, 3, 7]
    a2 = [4, 8, 3]
    diff = find_missing_items(a1, a2)

    # output: [6,7]
    print(diff)

    a1 = [4, 6, 8, 3]
    a2 = [4, 8, 3]
    diff = find_the_only_missing_item(a1, a2)
    # output: 6
    print(diff)

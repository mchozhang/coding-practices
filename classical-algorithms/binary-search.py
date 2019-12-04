# binary search
# find the first occurrence in the sorted array


def recursive_binary_search(arr, target, low, high):
    """
    binary search using recursive function,
    return -1 if target not found
    time complexity: O(logN)
    :param arr: sorted array
    :param target: target
    :param low: index of start position
    :param high: index of the end position
    :return: index of the first occurrence target
    """
    if low < high:
        mid = (high + low) // 2
        if target <= arr[mid]:
            return recursive_binary_search(arr, target, low, mid)
        else:
            return recursive_binary_search(arr, target, mid + 1, high)

    else:
        if arr[low] == target:
            return low
        else:
            return -1


def binary_search(arr, target):
    """
    binary search using non-recursive function
    :param arr: sorted array
    :param target: target
    :return: index of the first occurrence target
    """
    low = 0
    high = len(arr) - 1
    while low < high:
        mid = (low + high) // 2

        if arr[mid] >= target:
            high = mid
        else:
            low = mid + 1

    if arr[low] == target:
        return low
    else:
        return -1


if __name__ == "__main__":
    arr = [0, 1, 2, 4, 4, 4, 8, 8, 8, 9]

    # search 1, output 1
    result = recursive_binary_search(arr, 1, 0, len(arr) - 1)
    result2 = binary_search(arr, 1)
    print(result, result2)

    # search 4, output 3
    result = recursive_binary_search(arr, 4, 0, len(arr) - 1)
    result2 = binary_search(arr, 4)
    print(result, result2)

    # search 8, output 6
    result = recursive_binary_search(arr, 8, 0, len(arr) - 1)
    result2 = binary_search(arr, 8)
    print(result, result2)

    # search -1, output -1
    result = recursive_binary_search(arr, -1, 0, len(arr) - 1)
    result2 = binary_search(arr, -1)
    print(result, result2)

    # search 10, output -1
    result = recursive_binary_search(arr, 10, 0, len(arr) - 1)
    result2 = binary_search(arr, 10)
    print(result, result2)

    # search 5, output -1
    result = recursive_binary_search(arr, 5, 0, len(arr) - 1)
    result2 = binary_search(arr, 5)
    print(result, result2)



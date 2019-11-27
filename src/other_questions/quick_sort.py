# quick sort


def quick_sort(arr, low, high):
    """
    quick sort
    time complexity: O(NlogN)
    :param arr: array
    :param low: start index
    :param high: end index
    """
    if low < high:
        pivot = arr[high]

        j = low - 1  # smaller element index
        i = low  # loop index
        while i < high:
            if arr[i] < pivot:
                j += 1
                arr[i], arr[j] = arr[j], arr[i]
            i += 1

        # swap the pivot to the correct position
        arr[j + 1], arr[high] = arr[high], arr[j + 1]

        quick_sort(arr, low, j)
        quick_sort(arr, j + 1, high)


if __name__ == "__main__":
    arr = [6, 5, 7, 0, 1, 8, 2, 9, 3, 4]
    quick_sort(arr, 0, len(arr) - 1)
    print(arr)

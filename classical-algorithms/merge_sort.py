# merge sort


def merge_sort(arr):
    """
    merge sort
    time complexity: O(NlogN)
    :param arr: array
    """
    if len(arr) > 1:
        mid = len(arr) // 2
        left_array = arr[:mid]
        right_array = arr[mid:]

        merge_sort(left_array)
        merge_sort(right_array)

        # i, j is the index of the sub-array,
        # k is the index of the merged array
        i = j = k = 0

        # merge 2 sorted arrays, copy the data to the original array
        while i < len(left_array) and j < len(right_array):
            if left_array[i] < right_array[j]:
                arr[k] = left_array[i]
                i += 1
            else:
                arr[k] = right_array[j]
                j += 1
            k += 1

        # copy the data of the rest of the array
        while i < len(left_array):
            arr[k] = left_array[i]
            i += 1
            k += 1

        while j < len(right_array):
            arr[k] = right_array[j]
            j += 1
            k += 1


if __name__ == "__main__":
    arr = [6, 5, 7, 0, 1, 8, 2, 9, 3, 4]
    merge_sort(arr)
    print(arr)
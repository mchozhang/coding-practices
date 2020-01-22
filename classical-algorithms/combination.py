# combination algorithm


def combination(arr, n):
    """
    combination of selecting N items from the array
    """
    res = []

    def recursiveCombination(arr, step, selected, n):
        if len(selected) == n:
            res.append(selected.copy())
            return

        if step > len(arr) - 1:
            return

        selected.append(arr[step])
        recursiveCombination(arr, step + 1, selected, n)

        selected.pop()
        recursiveCombination(arr, step + 1, selected, n)

    recursiveCombination(arr, 0, [], n)

    return res


if __name__ == "__main__":
    arr = [1, 2, 3, 4, 5, 6, 7]

    print(combination(arr, 3))

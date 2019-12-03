# Dynamic Programming:
# Given K eggs and an N storey building, in the worst case, find the minimum time F (0 <= F <= N) that have to try to
# guarantee to find the floor i, that the egg won't break when threw from [0, i-1] floor, but will break when threw
# from [i, N] floor, note that the egg can be re-use if it doesn't break after throwing.
# Example:
# input: K = 1, N = 5 output: 5
# input: K = 2, N = 5 output: 3


def dropping_eggs(eggs, storey):
    """
    find the minimum time that have to try to address floor i,
    time complexity: O( K * N * N)
    space complexity: O(KN)
    :param eggs: number of egg
    :param storey: number of storey
    :return: result
    """
    result_record = dict()

    def recursive_computing(eggs, storey):
        if storey == 0:
            return 0

        if eggs == 1:
            return storey

        # check the result records
        if (eggs, storey) in result_record:
            return result_record[(eggs, storey)]
        else:
            # compute the result by dropping from every floor
            result = float('inf')
            for i in range(1, storey + 1):
                # drop the egg at floor i

                # the egg didn't break, search from the lower floors
                lower = 1 + recursive_computing(eggs, i - 1)

                # the egg broke, search from the upper floors
                upper = 1 + recursive_computing(eggs - 1, storey - i)

                result = min(result, max(lower, upper))

            # record the result
            result_record[(eggs, storey)] = result

            return result

    return recursive_computing(eggs, storey)


if __name__ == "__main__":
    res = dropping_eggs(2, 100)
    print(res)

#  6. ZigZag Conversion
#  https://leetcode.com/problems/zigzag-conversion/
#
#  submission : faster than 60%


def convert(s: str, numRows: int) -> str:
    if numRows == 1:
        return s

    zigzag = ["" for i in range(numRows)]

    dir = (1, 0)
    row, col = 0, 0
    for i, c in enumerate(s):
        zigzag[row] += c
        nextRow = row + dir[0]
        if nextRow < 0:
            dir = (1, 0)
        elif nextRow >= numRows:
            dir = (-1, 1)

        row += dir[0]
        col += dir[1]

    res = ""
    for row in zigzag:
        res += row
    return res


if __name__ == "__main__":
    a = "PAYPALISHIRING"

    # output "PAHNAPLSIIGYIR"
    print(convert(a, 3))

    # output "PINALSIGYAHRPI"
    print(convert(a, 4))

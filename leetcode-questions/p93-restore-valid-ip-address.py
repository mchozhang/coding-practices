#  93. Restore IP Addresses
#  https://leetcode.com/problems/restore-ip-addresses/
#
#  submission : faster than 94%

def restoreIpAddresses(s: str) -> List[str]:
    res = []

    def find(currentStr: str, remainStr: str, remain: int) -> None:
        if remain == 1:
            if validate(remainStr):
                res.append(currentStr + remainStr)
            return

        remain -= 1
        for i in range(1, 4):
            newStr, newRemainStr = remainStr[0:i], remainStr[i:]
            if remain <= len(newRemainStr) <= remain * 3 and validate(newStr):
                find(currentStr + newStr + ".", newRemainStr, remain)

    def validate(s: str) -> bool:
        return len(s) == 1 or (s[0] != '0' and (len(s) == 2 or (len(s) == 3 and s < "256")))

    find("", s, 4)
    return res

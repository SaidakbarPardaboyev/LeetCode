class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        n = 0
        res = 0
        for i in range(len(columnTitle)-1, -1, -1):
            res += ((ord(columnTitle[i]) - ord("A") + 1) * 26 ** n)
            n += 1
        return res
        
print(Solution().titleToNumber("FXSHRXW"))  # 2147483647
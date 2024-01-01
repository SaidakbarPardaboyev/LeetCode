class Solution:
    def thousandSeparator(self, n: int) -> str:
        n = str(n)
        if not len(n) > 3:
            return n

        n = n[::-1]

        i = 3
        while i < len(n):
            n = n[:i] + '.' + n[i:]
            i += 4

        return n[::-1]
    
print(Solution().thousandSeparator(1560211551))
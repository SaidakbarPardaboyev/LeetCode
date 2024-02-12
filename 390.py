class Solution:
    def lastRemaining(self, n: int) -> int:
        ls = [i for i in range(1, n+1)]
        lamp = True
        while len(ls) > 1:
            if lamp:
                i = 0
                while i < len(ls) and len(ls) > 1:
                    ls.pop(i)
                    i += 1
                lamp = False
            else:
                i = len(ls)-1
                while i >= 0 and len(ls) > 1:
                    ls.pop(i)
                    i -= 2
                lamp = True

        return ls[0]
    
print(Solution().lastRemaining(10))
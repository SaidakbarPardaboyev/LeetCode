class Solution:
    def countBits(self, n: int) -> list[int]:
        res = list()
        for i in range(0, n+1):
            res.append(bin(i)[2:].count('1'))

        return res
    
print(Solution().countBits(5))
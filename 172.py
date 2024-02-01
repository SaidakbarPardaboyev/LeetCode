class Solution:
    def trailingZeroes(self, n: int) -> int:
        res = 1
        for i in range(2, n+1):
            res *= i

        count = 0
        while res % 10 == 0:
            count += 1
            res //= 10

        return count
    

print(Solution().trailingZeroes(15))
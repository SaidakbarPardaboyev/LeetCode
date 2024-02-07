class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        from math import comb
        return comb(m+n-2, m-1)

# Example usage
solution = Solution()
print(solution.uniquePaths(23, 12))

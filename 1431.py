class Solution:
    def kidsWithCandies(self, candies: list[int], extraCandies: int) -> list[bool]:
        res = list()
        Max = max(candies)

        for i in candies:
            res.append(extraCandies + i >= Max)

        return res
    
candy = Solution()
print(candy.kidsWithCandies([2,3,5,1,3], extraCandies = 3))
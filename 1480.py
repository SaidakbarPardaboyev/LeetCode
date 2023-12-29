class Solution:
    def runningSum(self, nums: list[int]) -> list[int]:
        Summa = 0
        res = list()

        for i in nums:
            Summa += i
            res.append(Summa)

        return res
    
result = Solution()
print(result.runningSum([1,2,3,4]))
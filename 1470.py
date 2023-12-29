class Solution:
    def shuffle(self, nums: list[int], n: int) -> list[int]:
        res = list()

        for i in range(n):
            res.extend([nums[i], nums[n + i]])

        return res
    
result = Solution()
print(result.shuffle([2,5,1,3,4,7], 3))
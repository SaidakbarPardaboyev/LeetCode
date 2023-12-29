class Solution:
    def kLengthApart(self, nums: list[int], k: int) -> bool:
        sch = 0
        ind = 0
        while nums[ind] != 1:
            ind += 1
            if ind == len(nums):
                if len(nums) >= k:
                    return True
                else:
                    return False

        for i in range(ind + 1, len(nums)):
            if nums[i] == 1:
                if not sch >= k:
                    return False
                sch = 0
            else:
                sch += 1
        return True
    
result = Solution()
print(result.kLengthApart([1,0,0,0,1,0,0,1], 2))
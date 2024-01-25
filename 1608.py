class Solution:
    def specialArray(self, nums: list[int]) -> int:
        
        mx = -1
        for i in range(1, len(nums)+1):
            count = 0
            for j in nums:
                if i <= j:
                    count += 1
        
            if count == i:
                mx = i
            
        return mx
    
print(Solution().specialArray([3,5]))
class Solution:
    def threeConsecutiveOdds(self, arr: list[int]) -> bool:
        
        for i in range(2, len(arr)):
            if arr[i] % 2 != 0 and arr[i-1] % 2 != 0 and arr[i-2] % 2 != 0:
                return True
        
        return False
    
print(Solution().threeConsecutiveOdds([1,2,34,3,4,5,7,23,12]))
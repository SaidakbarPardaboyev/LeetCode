class Solution:
    def findKthPositive(self, arr: list[int], k: int) -> int:
        positiveNumber = 1
        while k != 0:
            if positiveNumber not in arr:
                k -= 1
            positiveNumber += 1
        
        return positiveNumber - 1
    
print(Solution().findKthPositive([2,3,4,7,11], 5))
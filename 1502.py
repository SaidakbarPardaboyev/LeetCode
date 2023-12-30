class Solution:
    def canMakeArithmeticProgression(self, arr: list[int]) -> bool:
        arr.sort()
        step = arr[1] - arr[0]
        for i in range(2, len(arr)):
            if arr[i] - arr[i-1] != step:
                return False

        return True
    
print(Solution().canMakeArithmeticProgression([3,5,1]))
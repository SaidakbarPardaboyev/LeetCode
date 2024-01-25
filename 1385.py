class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        res = 0
        for i in arr1:
            lamp = True
            for j in arr2:
                if abs(i-j) <= d:
                    lamp = False
            if lamp:
                res += 1
        return res
    

print(Solution().findTheDistanceValue([4,5,8], arr2 = [10,9,1,8], d = 2))
class Solution:
    def canBeEqual(self, target: list[int], arr: list[int]) -> bool:
        if len(arr) != len(target):
            return False
        tem = set(target)

        for i in tem:
            if target.count(i) != arr.count(i):
                return False
        
        return True
    
result = Solution()
print(result.canBeEqual([1,2,3,4], [2,4,1,3]))
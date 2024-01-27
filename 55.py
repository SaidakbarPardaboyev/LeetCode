class Solution:
    def canJump(self, nums: list[int]) -> bool:
        ls = list()
        res = False
        ls.append(True)
        for i in range(len(nums)-1, -1, -1):
            if nums[i] >= len(ls)-1:
                ls.append(True)
                res = True
            elif nums[i] == 0:
                ls.append(False)
                res = False
            else:
                num = nums[i]
                lamp = False
                for i in range(1, num+1):
                    if ls[-i]:
                        lamp = True
                        res = True
                        break
                if lamp:
                    ls.append(True)
                    res = True
                else:
                    ls.append(False)
                    res = False
        if res:
            return True 
        else:
            return False
        
print(Solution().canJump([2,3,1,1,4]))
class Solution:
    def intersection(self, nums1: list[int], nums2: list[int]) -> list[int]:
        ls = list()
        for i in set(nums1):
            if i in nums2:
                ls.append(i)
        
        return ls
    
print(Solution().intersection([1,2,2,1], nums2 = [2,2]))
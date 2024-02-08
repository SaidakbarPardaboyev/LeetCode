class Solution:
    def majorityElement(self, nums: list[int]) -> list[int]:
        res = list()
        leng = len(nums) // 3

        for i in set(nums):
            if nums.count(i) > leng:
                res.append(i)

        return res
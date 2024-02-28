class Solution:
    def isPossibleToSplit(self, nums: list[int]) -> bool:
        dic = dict()
        for i in nums:
            if i in dic.keys():
                dic[i] += 1
                if dic[i] > 2:
                    return False
            else:
                dic[i] = 1
        return True

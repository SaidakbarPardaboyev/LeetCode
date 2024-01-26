class Solution:
    def firstMissingPositive(self, nums: list[int]) -> int:
        nums.sort()
        finish = len(nums)-1

        num = 1
        while num <= finish+1:
            if not self.Search_i(num, nums, finish):
                return num
            num += 1

        return finish + 2

    def Search_i(self, i, nums, finish):
        start = 0

        while start <= finish:
            mid = (start + finish) // 2

            if i == nums[mid]:
                return True
            elif nums[mid] < i:
                start = mid + 1
            else:
                finish = mid - 1
        return False
    
print(Solution().firstMissingPositive([3,4,-1,1]))
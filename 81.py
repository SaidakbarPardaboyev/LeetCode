class Solution:
    def search(self, nums: list[int], target: int) -> bool:
        self.sorter(nums)
        start = 0
        finish = len(nums)-1

        while start <= finish:
            mid = (start + finish) // 2

            if nums[mid] == target:
                return True
            elif nums[mid] < target:
                start = mid + 1
            else:
                finish = mid - 1

        return False
    
    def sorter(self, nums):
        leng = len(nums)
        for i in range(leng):
            ind = i
            for j in range(i + 1, leng):
                if nums[ind] > nums[j]:
                    ind = j
            nums[ind], nums[i] = nums[i], nums[ind]

print(Solution().search([2,5,6,0,0,1,2], target = 0))
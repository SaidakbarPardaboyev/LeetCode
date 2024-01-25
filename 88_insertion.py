class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        nums1.extend(nums2)
        for i in range(n):
            nums1.remove(0)

        for i in range(1, len(nums1)):
            box = nums1[i]
            j = i-1
            while j >= 0 and nums1[j] > box:
                if nums1[j] > box:
                    nums1[j+1] = nums1[j]
                j -= 1
            nums1[j+1] = box
        print(nums1)

res = Solution()
res.merge([1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3)
        
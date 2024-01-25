class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        nums1.extend(nums2)
        for i in range(n):
            nums1.remove(0)

        leng = len(nums1)
        for i in range(leng):
            ind = i
            for j in range(i+1, leng):
                if nums1[j] < nums1[ind]:
                    ind = j
            nums1[i], nums1[ind] = nums1[ind], nums1[i]
        print(nums1)

res = Solution()
res.merge([1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3)
        
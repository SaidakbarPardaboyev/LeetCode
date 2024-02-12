class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        st = nums.copy()
        st = list(set(st))
        st.sort(key=lambda x: nums.count(x))
        return st[-k:]
    
print(Solution().topKFrequent([1,1,1,2,2,3], 2))
class Solution:

    def __init__(self, nums: List[int]):
        self.origin = nums.copy()
        self.ls = nums.copy()

    def reset(self) -> List[int]:
        self.ls = self.origin.copy()
        return self.ls

    def shuffle(self) -> List[int]:
        random.shuffle(self.ls)
        return self.ls


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()
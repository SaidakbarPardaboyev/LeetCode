def minSubsequence(nums: list[int]) -> list[int]:
    nums = sorted(nums)
    res = list()

    while not sum(res) > sum(nums):
        res.append(nums[-1])
        nums.pop()
    
    return res

print(minSubsequence([4,3,10,9,8]))
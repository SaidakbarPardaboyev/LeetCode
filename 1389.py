def createTargetArray(nums: list[int], index: list[int]) -> list[int]:
    res = list()
    for i in range(len(nums)):
        res.insert(index[i], nums[i])
    return res

print(createTargetArray([0,1,2,3,4], [0,1,2,2,1]))
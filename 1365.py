def smallerNumbersThanCurrent(nums: list[int]) -> list[int]:
    res = list()
    for i in nums:
        count = 0
        for j in nums:
            if i > j:
                count += 1
        res.append(count)
    return res

print(smallerNumbersThanCurrent([8,1,2,2,3]))
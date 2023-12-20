def findNumbers(nums: list[int]) -> int:
    res = 0

    for i in nums:
        i = str(i)
        if len(i) % 2 == 0:
            res += 1

    return res 

print(findNumbers([12,345,2,6,7896]))
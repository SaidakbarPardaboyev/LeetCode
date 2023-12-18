def maxProductDifference(nums: list[int]) -> int:
    max1 = max(nums)
    nums.remove(max1)
    max2 = max(nums)
    min1 = min(nums)
    nums.remove(min1)
    min2 = min(nums)
    return (max1 * max2) - (min1 * min2)

print(maxProductDifference([5,6,2,7,4]))
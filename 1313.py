def decompressRLElist(nums: list[int]) -> list[int]:
    ls = list()
    for i in range(0, len(nums), 2):
        ls.extend([nums[i + 1]] * nums[i])
    return ls

print(decompressRLElist([1,2,3,4]))
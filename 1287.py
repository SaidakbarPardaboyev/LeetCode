def findSpecialInteger(arr: list[int]) -> int:
    tem = set(arr)
    for i in tem:
        if arr.count(i) / len(arr) * 100 > 25:
            return i
    return 0

print(findSpecialInteger([1,2,2,6,6,6,6,7,10]))
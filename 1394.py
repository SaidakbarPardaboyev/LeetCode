def findLucky(arr: list[int]) -> int:
    res = -1
    st = set(arr)
    for i in st:
        if i == arr.count(i) and i > res:
            res = i
    
    return res

print(findLucky([1,2,2,3,3,3]))
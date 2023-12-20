def arrayRankTransform(arr: list[int]) -> list[int]:
    arrSet = list(set(arr))
    arrSet.sort()
    for i in range(len(arr)):
        arr[i] = arrSet.index(arr[i]) + 1
    return arr

print(arrayRankTransform([37,12,28,9,100,56,80,5,12]))

# [5,3,4,2,8,6,7,1,3]
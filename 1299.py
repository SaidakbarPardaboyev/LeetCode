def replaceElements(arr: list[int]) -> list[int]:
    ls = list()

    right_max = -1

    for i in range(len(arr) - 1, -1, -1):
        ls.append(right_max)
        right_max = max(right_max, arr[i])
    
    ls.reverse()
    return ls

print(replaceElements([17,18,5,4,6,1]))
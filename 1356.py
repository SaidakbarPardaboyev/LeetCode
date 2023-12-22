def sortByBits(arr: list[int]) -> list[int]:
    for i in range(len(arr)):
        arr[i] = [arr[i], bin(arr[i]).count('1')]
    
    for i in range(len(arr)):
        for j in range(i + 1, len(arr)):
            if arr[i][-1] > arr[j][-1]:
                arr[i], arr[j] = arr[j], arr[i]
            elif arr[i][-1] == arr[j][-1] and arr[i][0] > arr[j][0]:
                arr[i], arr[j] = arr[j], arr[i]

    for i in range(len(arr)):
        arr[i] = arr[i][0]

    return arr

print(sortByBits([0,1,2,5,3,4,5,6,7,8]))
def checkIfExist(arr: list[int]) -> bool:
    for i in arr:
        if i == i * i:
            if arr.count(i) >= 2:
                return True
        elif i * 2 in arr:
            return True
    return False

print(checkIfExist([10,2,5,3]))
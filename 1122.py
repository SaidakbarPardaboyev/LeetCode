def relativeSortArray(arr1: list[int], arr2: list[int]) -> list[int]:
    st = set(arr1).difference(set(arr2))
    st = sorted(list(st))

    res = list()
    for i in arr2:
        sch = arr1.count(i)
        for j in range(sch):
            res.append(i)
    for i in st:
        sch = arr1.count(i)
        for j in range(sch):
            res.append(i)
    return res

print(relativeSortArray([2,21,43,38,0,42,33,7,24,13,12,27,12,24,5,23,29,48,30,31], [2,42,38,0,43,21]))

    # [2,42,38,0,43,21,5,7,12,12,13,23,24,24,27,29,30,31,33,48]
def uniqueOccurrences(arr: list[int]) -> bool:
    tem = set(arr)

    ls = list()

    for i in tem:
        ls.append(arr.count(i))

    return len(ls) == len(set(ls))

print(uniqueOccurrences([1,2,2,1,1,3]))
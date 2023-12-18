def countCharacters(words: list[str], chars: str) -> int:
    res = 0
    for i in words:
        st = set(i)
        sch = 0
        for j in st:
            if not i.count(j) <= chars.count(j):
                sch += 1
                break
        if sch == 0:
            res += len(i)
    return res


print(countCharacters(["cat","bt","hat","tree"], "atach"))
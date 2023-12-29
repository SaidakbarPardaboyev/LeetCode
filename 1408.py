def stringMatching(words: list[str]) -> list[str]:
    res = set()

    for i in range(len(words)):
        sch = 0
        for j in range(0, len(words)):
            if words[i] in words[j]:
                sch += 1
        if sch > 1:
            res.add(words[i])
    return res

print(stringMatching(["mass","as","hero","superhero"]))
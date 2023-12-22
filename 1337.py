def kWeakestRows(mat: list[list[int]], k: int) -> list[int]:
    for i in range(len(mat)):
        mat[i] = mat[i].count(1)
    ls = list()
    for i in range(k):
        ls.append(mat.index(min(mat)))
        mat[mat.index(min(mat))] = max(mat) + 1
    return ls

print(kWeakestRows([[1,1,0,0,0],
                    [1,1,1,1,0],
                    [1,0,0,0,0],
                    [1,1,0,0,0],
                    [1,1,1,1,1]], 3))
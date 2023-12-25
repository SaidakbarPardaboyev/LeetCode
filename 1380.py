def luckyNumbers (matrix: list[list[int]]) -> list[int]:
    res = list()

    ls = [max(col) for col in zip(*matrix)]

    for i in range(len(matrix)):
        tem = min(matrix[i])
        temls = matrix[i]
        ind = temls.index(tem)
        if matrix[i][ind] == ls[ind]:
            res.append(tem)

    return res

print(luckyNumbers([[3,7,8],
                    [9,11,13],
                    [15,16,17]]))   
def shiftGrid(grid: list[list[int]], k: int) -> list[list[int]]:
    ls = list()

    for i in range(len(grid)):
        for j in range(len(grid[i])):
            ls.append(grid[i][j])
    for i in range(k):
        ls.insert(0, ls[-1])
        ls.pop(-1)

    ls2 = list()
    k = 0
    for i in range(len(grid)):
        tem = list()
        for j in range(len(grid[i])):
            tem.append(ls[k])
            k += 1    
        ls2.append(tem)
    return ls2

print(shiftGrid([[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], 4))
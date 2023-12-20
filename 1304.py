def sumZero(n: int) -> list[int]:
    ls = list()
    if n % 2 != 0:
        ls.append(0)
    
    for i in range(1, n // 2 + 1):
        ls.append(i)
        ls.append(i * -1)
    return ls

print(sumZero(6))
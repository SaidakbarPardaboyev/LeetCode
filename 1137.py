def tribonacci(n: int) -> int:
    T0 = 0
    T1 = 1
    T2 = 1
    if n == 0:
        return T0
    elif n == 1:
        return T1
    elif n == 2:
        return T2
    for i in range(n - 2):
        T3 = T0 + T1 + T2
        T0 = T1
        T1 = T2
        T2 = T3
    return T2


print(tribonacci(4))
def numberOfSteps(num: int) -> int:
    res = 0
    while num != 0:
        if num % 2 != 0 and num != 0:
            num -= 1
            res += 1
        num //= 2
        res += 1
    if res != 0:
        res -= 1
    return res

print(numberOfSteps(14))
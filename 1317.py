def getNoZeroIntegers(n: int) -> list[int]:
    for i in range(1, n):
        if str(i).count('0') == 0 and str(n - i).count('0') == 0:
            return [i, n - i]

print(getNoZeroIntegers(11))
def subtractProductAndSum(n: int) -> int:
    tem = str(n)
    tem = list(tem)

    mult = 1
    summ = 0

    for i in tem:
        mult *= int(i)
        summ += int(i)

    return mult - summ

print(subtractProductAndSum(234))
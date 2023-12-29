def reformat(s: str) -> str:
    digit = list()
    alpha = list()
    schD = 0
    schA = 0
    for i in s:
        if i.isdigit():
            schD += 1
            digit.append(i)
        elif i.isalpha():
            schA += 1
            alpha.append(i)
    ls = [-1, 1, 0]
    j = 0
    if schD - schA not in ls:
        return ""
    elif schD - schA == -1:
        for i in range(len(alpha)):
            digit.insert(j, alpha[i])
            j += 2
        return digit
    else:
        for i in range(len(digit)):
            alpha.insert(j, digit[i])
            j += 2
        return alpha

print(reformat("a0b1c2"))
print(reformat("leetcode"))
print(reformat("1229857369"))
    
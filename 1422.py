def maxScore(s: str) -> int:
    res = 0

    for i in range(1, len(s)):
        tem = s[:i].count('0') + s[i:].count('1')
        if tem > res:
            res = tem

    return res

print(maxScore("011101"))
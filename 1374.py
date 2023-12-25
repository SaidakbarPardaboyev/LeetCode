def generateTheString(n: int) -> str:
    st = str()
    if n == 2:
        st = "ab"
    if n % 2 == 0:
        if (n // 2) % 2 == 0:
            st = ('a' * (n // 2 - 1)) + ('b' * (n // 2 + 1))
        else:
            st = ('a' * (n // 2)) + ('b' * (n // 2))
    else:
        st = ('a' * n)
    return st

print(generateTheString(2))
def freqAlphabets(s: str) -> str:
    ls = list()
    i = 0
    while len(s) > 0:
        if s.count("#") != 0 and s[i + 2] == "#":
            ls.append(chr(int(s[i:i + 2]) + ord('a') - 1))
            s = s[i + 3:]
        else:
            ls.append(chr(int(s[i]) + ord('a') - 1))
            s = s[i + 1:]
    
    return "".join(ls)

print(freqAlphabets("1326#"))
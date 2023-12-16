def isAnagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    tem = set(s)
    for i in tem:
        if s.count(i) != t.count(i):
            return False
    return True

print(isAnagram("anagram", "nagaram"))
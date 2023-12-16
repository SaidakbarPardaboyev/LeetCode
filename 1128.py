def numEquivDominoPairs(dominoes: list[list[int]]) -> int:
        count = {}
        res = 0
        
        for domino in dominoes:
            domino.sort()
            key = tuple(domino)
            count[key] = count.get(key, 0) + 1
        
        for val in count.values():
            res += val * (val - 1) // 2  # Using the combination formula to count pairs
        
        return res

print(numEquivDominoPairs([[1,2],[2,1],[3,4],[5,6]]))
print(numEquivDominoPairs([[1,2],[1,2],[1,1],[1,2],[2,2]]))

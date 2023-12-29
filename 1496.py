class Solution:
    def isPathCrossing(self, path: str) -> bool:
        res = list()
        Cor = [0, 0]
        res.append(tuple(Cor))

        for i in path:
            if i == 'N':
                Cor[0] -= 1
            elif i == 'S':
                Cor[0] += 1
            elif i == 'W':
                Cor[1] -= 1
            elif i == 'E':
                Cor[1] += 1

            res.append(tuple(Cor))
        
        return len(res) != len(set(res))
        
print(Solution().isPathCrossing("NESWW"))
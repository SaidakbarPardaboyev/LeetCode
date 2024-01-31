class Solution:
    def findCenter(self, edges: list[list[int]]) -> int:
        for k in edges[0]:
            sch = 0
            for i in edges:
                if k in i:
                    sch += 1
            if sch == len(edges):
                return k
        
print(Solution().findCenter([[1,2],[5,1],[1,3],[1,4]]))
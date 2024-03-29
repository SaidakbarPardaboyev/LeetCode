class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        v1 = version1.split('.')
        v2 = version2.split('.')

        for i in range(len(v1)):
            v1[i] = int(v1[i])
        
        for i in range(len(v2)):
            v2[i] = int(v2[i])

        leng = min(len(v1), len(v2))

        for i in range(leng):
            if int(v1[i]) > int(v2[i]):
                return 1
            elif int(v1[i]) < int(v2[i]):
                return -1
        if len(v1) > len(v2):
            for i in range(len(v2),len(v1)):
                if v1[i] != 0:
                    return 1
            return 0
        elif len(v1) < len(v2):
            for i in range(len(v1),len(v2)):
                if v2[i] != 0:
                    return -1
            return 0
        return 0

print(Solution().compareVersion("1.1", "1.001.0"))
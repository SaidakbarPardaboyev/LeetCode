class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        res = -1
        sentence = sentence.split()

        for i in range(len(sentence)):
            if sentence[i].startswith(searchWord):
                res = i + 1
                break
        
        return res
    
result = Solution()
print(result.isPrefixOfWord("i love eating burger", "burg"))
class Solution:
    def flipAndInvertImage(self, image: List[List[int]]) -> List[List[int]]:
        # revesing digits of each row
        for k in range(len(image)):
            j = len(image[k])-1
            i = 0
            while i < j:
                image[k][i], image[k][j] = image[k][j], image[k][i]
                j -= 1
                i += 1
        
        # reverse each digit if it 1 to 0 or teskarisi
        for i in range(len(image)):
            for j in range(len(image[k])):
                if image[i][j]:
                    image[i][j] = 0
                else:
                    image[i][j] = 1
        
        return image
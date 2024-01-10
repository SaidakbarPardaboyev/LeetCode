class Solution:
    def reorderSpaces(self, text: str) -> str:
        ls = list()
        tem = str()
        sch = 0
        size = len(text)
        for i in range(len(text)):
            if text[i].isspace():
                sch += 1
                if len(tem) > 0:
                    ls.append(tem)
                tem = ""
            else:
                tem += text[i]
                if i == size - 1:
                    if len(tem) > 0:
                        ls.append(tem)
        if len(ls) > 1:            
            spaces = ' ' * (sch // (len(ls)-1))
            ls = f"{spaces}".join(ls) + ((sch % (len(ls)-1)) * ' ')
        else:
            ls = "".join(ls) + (sch * ' ')
        return ls


print(Solution().reorderSpaces(" practice   makes   perfect"))
class MinStack:

    def __init__(self):
        self.stack = list()
        self.min = None

    def push(self, val: int) -> None:
        self.stack.append(val)

    def pop(self) -> None:
        self.stack.pop()        

    def top(self) -> int:
        return self.stack[-1]
        
    def getMin(self) -> int:
        ls = self.stack.copy()
        return min(ls)


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
    

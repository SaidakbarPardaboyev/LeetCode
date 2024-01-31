import queue

class MyQueue:

    def __init__(self):
        self.queue = queue.Queue()

    def push(self, x: int) -> None:
        self.queue.put(x)

    def pop(self) -> int:
        if not self.queue.empty():
            return self.queue.get()
        return None

    def peek(self) -> int:
        if not self.queue.empty():
            tem = self.queue.queue[0]
            return tem
        else:
            return None

    def empty(self) -> bool:
        return self.queue.empty()


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
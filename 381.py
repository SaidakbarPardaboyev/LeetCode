class RandomizedCollection:

    def __init__(self):
        self.st = list()

    def insert(self, val: int) -> bool:
        if val not in self.st:
            self.st.append(val)
            return True
        self.st.append(val)
        return False

    def remove(self, val: int) -> bool:
        if val in self.st:
            self.st.remove(val)
            return True
        return False

    def getRandom(self) -> int:
        return random.choice(self.st)        

# Your RandomizedCollection object will be instantiated and called as such:
# obj = RandomizedCollection()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
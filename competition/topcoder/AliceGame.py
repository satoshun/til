class AliceGame:
    def findMinimumValue(self, x, y):
        s = x + y
        start = 1
        total = 0
        vs = []
        while True:
            total += start
            vs.append(start)
            if total == s:
                break
            elif total > s:
                return -1
            start += 2
        vs = vs[::-1]
        result = []
        self.padding(vs, x, result)
        return len(result)

    def padding(self, vs, size, result):
        for i in vs:
            if size == i:
                result.append(i)
                return result
            if size > i:
                size = size - i
                result.append(i)
                vs.remove(i)
                return self.padding(vs, size, result)

        return False


a = AliceGame()
assert a.findMinimumValue(8, 17) == 2
assert a.findMinimumValue(17, 8) == 3

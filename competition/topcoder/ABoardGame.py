class ABoardGame:
    def whoWins(self, board):
        j = ''.join(board)
        start = len(j) // 2 - 2
        l = 2
        while True:
            t = j[start:start+l*l]
            a = t.count('A')
            b = t.count('B')
            print(t)
            print(start)
            if a > b:
                return 'Alice'
            elif a < b:
                return 'Bob'
            l += 2
            start -= (l*l- (l-2)**2) // 2
            if start < 0:
                break
        return 'draw'

board = ABoardGame()
# assert board.whoWins([".....A", "......", "..A...", "...B..", "......", "......"]) == 'Alice'
assert board.whoWins(["AAAA", "A.BA", "A..A", "AAAA"]) == 'Bob'

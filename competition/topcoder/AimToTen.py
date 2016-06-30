class AimToTen:
    def need(self, marks):
        marks = list(marks)
        i = 0
        while True:
            if sum(marks) / float((len(marks))) >= 9.5:
                return i
            i += 1
            marks.append(10)


l = AimToTen()
assert l.need((8, 9)) == 4
assert l.need(0 for _ in range(50)) == 950

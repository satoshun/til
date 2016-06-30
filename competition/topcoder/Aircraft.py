import math

class Aircraft:
    def nearMiss(self, p1, v1, p2, v2, R):
        p1, v1, p2, v2 = list(p1), list(v1), list(p2), list(v2)
        before = self.dis(p1, p2)
        d = 5.0
        if before <= R:
            return 'YES'

        while True:
            p1[0] += v1[0] * d
            p1[1] += v1[1] * d
            p1[2] += v1[2] * d

            p2[0] += v2[0] * d
            p2[1] += v2[1] * d
            p2[2] += v2[2] * d

            cur = self.dis(p1, p2)
            if cur <= R:
                return 'YES'

            if before < cur:
                if d > 0:
                    d -= 0.25
                else:
                    d += 0.25
                d *= -1

            if abs(d) < 0.01:
                return 'NO'
            before = cur

        return 'NO'


    def dis(self, p1, p2):
        return math.sqrt((p1[0] - p2[0])**2 + (p1[1] - p2[1])**2 + (p1[2] - p2[2])**2)




l = Aircraft()
assert l.nearMiss((15, 50, 5), (25, 1, 0), (161, 102, 9), (-10, -10, -1), 10) == 'YES'
assert l.nearMiss((0, 0, 0), (2, 2, 0), (9, 0, 5), (-2, 2, 0), 5) == 'YES'
assert l.nearMiss((0, 0, 0), (-2, 2, 0), (9, 0, 5), (2, 2, 0), 5) == 'NO'

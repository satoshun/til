import itertools


class AkariDaisukiDiv2:
    def countTuples(self, S):
        cc = 0
        for i in range(1, len(S)):
            cc += self._c(S, i)
        return cc

    def _c(self, S, index):
        s = set()
        for i in range(len(S)-index):
            s.add(S[i:i+index])
        cc = 0
        for ss in s:
            c = S[1:-1].count(ss)
            if c >= 2:
                l = len([_ for _ in itertools.combinations('a' * c, 2)])
                cc += l

        return cc


l = AkariDaisukiDiv2()
# assert l.countTuples('topcoderdivtwo') == 2
# assert l.countTuples('foxciel') == 0
# assert l.countTuples('magicalgirl') == 4
assert l.countTuples('waaiusushioakariusushiodaisuki') == 75

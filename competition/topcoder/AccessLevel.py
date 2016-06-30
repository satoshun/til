class AccessLevel:
    def canAccess(self, rights, minPermission):
        d = []
        for r in rights:
            if r >= minPermission:
                d.append('A')
            else:
                d.append('D')
        return ''.join(d)


l = AccessLevel()
assert l.canAccess((0, 1, 2, 3, 4, 5), 2) == 'DDAAAA'

def g():
    yield range(10)
    yield from range(5)
    yield from range(5, 10)

print(g())
print(next(g()))
print(range(4))

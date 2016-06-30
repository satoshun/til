import asyncio

async def compute(x, y):
    print("Compute %s + %s ..." % (x, y))
    await asyncio.sleep(1.0)
    return x + y

async def print_sum(x, y):
    result = await compute(x, y)
    print("%s + %s = %s" % (x, y, result))
    return result

loop = asyncio.get_event_loop()
f = print_sum(1, 2)
loop.run_until_complete(f)
print(10000000000)
loop.close()
print(f.cr_code)

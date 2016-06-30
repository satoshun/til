import asyncio
import time

a = asyncio.sleep(5)
next(a)

# future = asyncio.Future()
# future.add_done_callback(lambda: print(100000))
# time.sleep(5)

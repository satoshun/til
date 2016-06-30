import threading
import time


# print(dir(threading._start_new_thread))
# print(threading.active_count())
# def hoge(i):
#     print(threading.active_count())
#     for j in range(1000000000):
#         pass
#     print(i)
# for i in range(5):
#     threading._start_new_thread(hoge, (i,))
# time.sleep(5)


class A:
    def __call__(self, *args, **kwargs):
        print("__call__", type(args), type(kwargs))

a = A()
a(1, 2, 3)

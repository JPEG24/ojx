from random import randint
from sortedcontainers import SortedSet

def ABC244_C():
	N = randint(1, 10)
	print(N, flush=True)

	st = SortedSet(i+1 for i in range(N*2+1))
	while True:
		x = int(input())
		assert x in st
		st.remove(x)

		if len(st) == 0:
			print(0, flush=True)
			return

		y = st[0]
		print(y, flush=True)
		st.remove(y)

ABC244_C()
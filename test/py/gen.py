from random import randint

# test random-test
def ABC451_C():
	Q = randint(1, 10)
	print(Q)
	for _ in range(Q):
		t = randint(1, 2)
		h = randint(1, 100)
		print(t, h)

ABC451_C()
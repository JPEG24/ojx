from random import randint
from sortedcontainers import SortedList, SortedSet

# test interactive problem
def ABC244_C_AC():
  N = int(input())

  st = SortedSet(i+1 for i in range(N*2+1))

  while True:
    x = st[0]
    print(x, flush=True)
    st.remove(x)

    y = int(input())
    if y == 0:
      return

    st.remove(y)


def ABC244_C_WA():
  N = int(input())

  while True:
    x = randint(1, N*2+1)
    print(x, flush=True)

    y = int(input())
    if y == 0:
      return


# test random_test
def ABC451_C_AC():
  Q = int(input())

  st = SortedList()

  for _ in range(Q):
    t, h = map(int, input().split())

    if t == 1:
      st.add(h)
    else:
      idx = st.bisect_right(h)
      del st[:idx]

    print(len(st))


def ABC451_C_WA():
  Q = int(input())

  st = SortedSet()

  for _ in range(Q):
    t, h = map(int, input().split())

    if t == 1:
      st.add(h)
    else:
      idx = st.bisect_right(h)
      del st[:idx]

    print(len(st))


# test interactive problem: ABC244-C (Yamanote Line Game)
# ABC244_C_AC()
# ABC244_C_WA()

# test random_test: ABC451-C (Understory)
# ABC451_C_AC()
# ABC451_C_WA()
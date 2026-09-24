from sortedcontainers import SortedList

def ABC451_C():
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

ABC451_C()
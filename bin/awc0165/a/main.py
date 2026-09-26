n, q = map(int, input().split())
a = list(map(int, input().split()))
b = list(map(int, input().split()))
c = list(map(int, input().split()))

for i in range(q):
	x = min(a[c[i]-1], b[c[i]-1])
	if c[i] < n:
		a[c[i]] += a[c[i]-1]-x
	a[c[i]-1] = 0

print(*a)
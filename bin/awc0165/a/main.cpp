#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;
#define rep(i,n) for (int i = 0; i < (n); i++)

int main() {
	cin.tie(nullptr)->sync_with_stdio(false);
	cout << fixed << setprecision(15);

	int n, q;
	cin >> n >> q;
	vector<int> a(n), b(n), c(n);
	rep(i,n) cin >> a[i];
	rep(i,n) cin >> b[i];
	
	rep(_,q) {
		int c;
		cin >> c;
		c--;
		a[c] -= min(a[c], b[c]);
		if (c+1 < n) {
			a[c+1] += a[c];
		}
		a[c] = 0;
	}

	rep(i,n) cout << a[i] << " \n"[i == n-1];
	return 0;
}
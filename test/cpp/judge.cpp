#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;
#define rep(i,n) for(int i = 0; i < (n); ++i)

int randint(int l, int r) {
	static mt19937 mt(random_device{}());
	uniform_int_distribution<int> dist(l, r);
	return dist(mt);
}

// test interactive problem
void ABC244_C() {
	int N = randint(1,10);
	cout << N << endl;
	set<int> st;
	rep(i,N*2+1) st.insert(i+1);
	
	while (1) {
		int x;
		cin >> x;
		assert(st.count(x));
		st.erase(x);

		if (st.empty()) {
			cout << 0 << endl;
			return;
		}
		int y = *st.begin();
		cout << y << endl;
		st.erase(y);
	}
}

int main() {
	cin.tie(nullptr) -> sync_with_stdio(false);
	cout << fixed << setprecision(15);

	ABC244_C();
	return 0;
}
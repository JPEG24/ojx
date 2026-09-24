#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;

// test random test
void ABC451_C() {
	int Q;
	cin >> Q;
	multiset<int> st;
	while (Q--) {
		int t, h;
		cin >> t >> h;
		if (t == 1) {
			st.insert(h);
		} else {
			while (!st.empty() && *st.begin() <= h) {
				st.erase(st.begin());
			}
		}
		cout << st.size() << '\n';
	}
}

int main() {
	cin.tie(nullptr) -> sync_with_stdio(false);
	cout << fixed << setprecision(15);

	ABC451_C();

	return 0;
}
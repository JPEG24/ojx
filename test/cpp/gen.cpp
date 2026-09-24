#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;

int randint(int l, int r) {
	static mt19937 mt(random_device{}());
	uniform_int_distribution<int> dist(l, r);
	return dist(mt);
}

// test random test
void ABC451_C() {
	int Q;
	Q = randint(1,10);
	cout << Q << endl;
	while (Q--) {
		int t = randint(1,2);
		int h = randint(1,10);
		cout << t << ' ' << h << endl;
	}
}

int main() {
	cin.tie(nullptr) -> sync_with_stdio(false);
	cout << fixed << setprecision(15);

	ABC451_C();
	return 0;
}
#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;
#define rep(i,n) for(int i = 0; i < (n); ++i)

// test interactive problem
void ABC244_C_AC() {
  int N;
  cin >> N;
  set<int> st;
  rep(i,N*2+1) st.insert(i+1);

  while (1) {
    int x = *st.begin();
    // int x = 1;
    cout << x << endl;
    st.erase(x);

    int y;
    cin >> y;
    if (y == 0) return;
    st.erase(y);
  }
}

void ABC244_C_WA() {
  int N;
  cin >> N;
  
  while (1) {
    int x = 1;
    cout << x << endl;

    int y;
    cin >> y;
    if (y == 0) return;
  }
}

// test random_test
void ABC451_C_AC() {
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

void ABC451_C_WA() {
  int Q;
  cin >> Q;
  set<int> st;
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

  // test interactive problem: ABC244-C (Yamanote Line Game)
  // ABC244_C_AC();
  // ABC244_C_WA();

  // test random_test: ABC451-C (Understory)
  // ABC451_C_AC();
  // ABC451_C_WA();
  return 0;
}
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

typedef long long ll;

const int MAX_VAL = 1000005;
ll tree[4 * MAX_VAL];

void update(int node, int start, int end, int idx, ll val) {
    if (start == end) {
        tree[node] = max(tree[node], val);
        return;
    }
    int mid = (start + end) / 2;
    if (idx <= mid) update(2 * node, start, mid, idx, val);
    else update(2 * node + 1, mid + 1, end, idx, val);
    tree[node] = max(tree[2 * node], tree[2 * node + 1]);
}

ll query(int node, int start, int end, int l, int r) {
    if (r < start || end < l) return 0;
    if (l <= start && end <= r) return tree[node];
    int mid = (start + end) / 2;
    return max(query(2 * node, start, mid, l, r), query(2 * node + 1, mid + 1, end, l, r));
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    int n, k;
    cin >> n >> k;

    ll max_total_weight = 0;
    for (int i = 0; i < n; ++i) {
        int v;
        ll w;
        cin >> v >> w;

        int l = max(0, v - k);
        int r = min(MAX_VAL - 1, v + k);

        ll best_prev = query(1, 0, MAX_VAL - 1, l, r);
        ll current_dp = best_prev + w;

        update(1, 0, MAX_VAL - 1, v, current_dp);
        max_total_weight = max(max_total_weight, current_dp);
    }

    cout << max_total_weight << endl;

    return 0;
}
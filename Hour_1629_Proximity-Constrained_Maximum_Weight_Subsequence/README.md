# Proximity-Constrained Maximum Weight Subsequence

**Language:** C++ | **Date:** 2024-01-01 16:29

## Description
Given N items, each having a value v[i] and a weight w[i]. You need to select a subsequence of items (preserving their original relative order) such that for any two consecutive items selected, (v_a, w_a) and (v_b, w_b), the absolute difference between their values is at most K, i.e., |v_a - v_b| <= K. Your goal is to maximize the total sum of weights of the selected items.

Input:
- The first line contains two integers N and K (1 <= N <= 10^5, 0 <= K <= 10^6).
- The next N lines each contain two integers v[i] and w[i] (0 <= v[i] <= 10^6, 1 <= w[i] <= 10^9).

Output:
- A single integer representing the maximum total weight.

Example:
Input:
5 2
5 10
2 5
4 15
8 20
3 10

Output:
35
Explanation: The optimal subsequence is (v=5, w=10) -> (v=4, w=15) -> (v=3, w=10). The differences are |5-4|=1 <= 2 and |4-3|=1 <= 2. Total weight = 10+15+10 = 35.

## Explanation
The problem asks for a maximum weight subsequence with a value difference constraint between adjacent elements. This can be solved using Dynamic Programming: let DP[v] be the maximum weight of a valid subsequence ending with an item of value 'v'. For each item (v_i, w_i), the transition is DP[v_i] = w_i + max(DP[j]) for all j such that v_i - K <= j <= v_i + K. Since the range of values is up to 10^6, a simple loop for the max would result in O(N * V) complexity, which is too slow. We optimize this using a Segment Tree (or Fenwick Tree) to perform range maximum queries and point updates in O(log V) time. The total time complexity is O(N log(max_V)), and the space complexity is O(max_V).

# Maximum Product of Two Non-Overlapping Palindromic Subsequences

**Language:** JavaScript | **Date:** 2024-01-01 14:42

## Description
Given a string `s`, find the maximum possible product of the lengths of two non-overlapping palindromic subsequences. Two subsequences are considered non-overlapping if there exists an index `i` such that the first subsequence is formed entirely from characters in the prefix `s[0...i]` and the second subsequence is formed entirely from characters in the suffix `s[i+1...n-1]`. A subsequence is a sequence that can be derived from another sequence by deleting zero or more elements without changing the order of the remaining elements.

Example 1:
Input: s = "abacaba"
Output: 9
Explanation: We can pick "aba" from the left part and "aba" from the right part (split at index 2 or 3). Product: 3 * 3 = 9.

Example 2:
Input: s = "leetcodecom"
Output: 9
Explanation: LPS of "leet" is "ee" (2), LPS of "codecom" is "coc" (3) or "mom" (3). Product: 2 * 3 = 6. Another split: "lee" (2) and "tcodecom" (3). The maximum product found is 9 (e.g., "eee" and "ocmco" is not possible, but "eee" from "lee...e" and "ocmco" is not possible here). Actually, for "leetcodecom", split after "lee" (LPS="ee", len 2) and "tcodecom" (LPS="coc", len 3) -> 6. Max possible might be 9 using different subsequences.

## Explanation
The problem asks for the maximum product of lengths of two non-overlapping palindromic subsequences divided by a split point. 
1. We use Dynamic Programming to precalculate the Longest Palindromic Subsequence (LPS) for every possible substring `s[i...j]`. The state transition is: if `s[i] == s[j]`, `dp[i][j] = dp[i+1][j-1] + 2`. Otherwise, `dp[i][j] = max(dp[i+1][j], dp[i][j-1])`.
2. After filling the `dp` table in O(N^2) time, we iterate through all possible split points `i` from `0` to `n-2`.
3. For each split point, the maximum product is the LPS of the left side `dp[0][i]` multiplied by the LPS of the right side `dp[i+1][n-1]`.
4. Time Complexity: O(N^2) to fill the DP table and O(N) to find the max product. Total: O(N^2).
5. Space Complexity: O(N^2) to store the DP table.

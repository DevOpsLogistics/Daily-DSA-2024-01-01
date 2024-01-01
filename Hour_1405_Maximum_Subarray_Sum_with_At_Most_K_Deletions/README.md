# Maximum Subarray Sum with At Most K Deletions

**Language:** Go | **Date:** 2024-01-01 14:05

## Description
Given an array of integers `nums` and an integer `k`, find the maximum possible sum of a non-empty subarray after deleting at most `k` elements from it. The remaining elements in the subarray must have been part of the same contiguous segment in the original array before the deletions. For example, in `nums = [1, -2, 0, 3]`, if we delete `-2`, the remaining elements are `[1, 0, 3]` which form a sum of 4. Note that you cannot delete all elements from a subarray; it must remain non-empty.

## Explanation
The problem is an extension of the Maximum Subarray Sum (Kadane's Algorithm). We use Dynamic Programming to track the maximum sum ending at each index with a specific number of deletions. Two states are maintained: 'keep[i][j]' (max sum including element i with j deletions) and 'skip[i][j]' (max sum excluding element i with j deletions). 
1. keep[i][j]: Can be starting a new subarray at i (if j=0) or adding nums[i] to a previous subarray that already had j deletions (either ending in a keep or a skip at i-1).
2. skip[i][j]: Can only be achieved by taking a subarray that ended at i-1 with j-1 deletions and choosing to ignore nums[i].
Time Complexity: O(N * K), where N is the number of elements and K is the allowed deletions. 
Space Complexity: O(N * K) to store the DP table, which can be optimized to O(K) by only keeping the previous row.

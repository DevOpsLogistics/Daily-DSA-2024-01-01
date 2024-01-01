package main

import (
	"fmt"
)

// max is a helper function to return the greater of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSumWithKDeletions(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	const INF = 1e16

	// keep[i][j] is the max sum of a subarray ending at index i with exactly j deletions, where nums[i] is included (kept).
	keep := make([][]int, n)
	// skip[i][j] is the max sum of a subarray ending at index i with exactly j deletions, where nums[i] is excluded (skipped).
	skip := make([][]int, n)

	for i := range keep {
		keep[i] = make([]int, k+1)
		skip[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			keep[i][j] = -INF
			skip[i][j] = -INF
		}
	}

	maxSum := -INF

	for i := 0; i < n; i++ {
		// Base case for j=0: standard Kadane's algorithm
		keep[i][0] = nums[i]
		if i > 0 {
			keep[i][0] = max(keep[i][0], keep[i-1][0]+nums[i])
		}
		maxSum = max(maxSum, keep[i][0])

		for j := 1; j <= k; j++ {
			// To compute keep[i][j], we must include nums[i].
			// We either start a new subarray (only valid if j=0, handled above) 
			// or extend from a previous state at i-1 with j deletions.
			if i > 0 {
				prevSum := max(keep[i-1][j], skip[i-1][j])
				if prevSum != -INF {
					keep[i][j] = max(nums[i], prevSum+nums[i])
				} else {
					keep[i][j] = nums[i]
				}
			} else {
				// At index 0, we can't have j > 0 deletions within a subarray ending at 0.
				keep[i][j] = nums[i]
			}

			// To compute skip[i][j], we exclude nums[i]. 
			// This must extend from a state at i-1 with j-1 deletions.
			if i > 0 {
				skip[i][j] = max(keep[i-1][j-1], skip[i-1][j-1])
			}
			
			// A subarray ending in a skip is technically the same as a subarray ending earlier.
			// So we only update maxSum with keep[i][j].
			maxSum = max(maxSum, keep[i][j])
		}
	}

	return maxSum
}

func main() {
	// Example 1
	nums1 := []int{1, -2, 0, 3}
	k1 := 1
	fmt.Printf("Input: %v, k: %d -> Output: %d\n", nums1, k1, maximumSumWithKDeletions(nums1, k1))

	// Example 2
	nums2 := []int{1, -2, -3, 4}
	k2 := 1
	fmt.Printf("Input: %v, k: %d -> Output: %d\n", nums2, k2, maximumSumWithKDeletions(nums2, k2))

	// Example 3 (All negative numbers)
	nums3 := []int{-10, -5, -2}
	k3 := 2
	fmt.Printf("Input: %v, k: %d -> Output: %d\n", nums3, k3, maximumSumWithKDeletions(nums3, k3))
}
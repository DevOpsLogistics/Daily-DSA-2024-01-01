/**
 * @param {string} s
 * @return {number}
 */
function maxProduct(s) {
    const n = s.length;
    if (n < 2) return 0;

    // dp[i][j] will store the length of the Longest Palindromic Subsequence in s[i...j]
    const dp = Array.from({ length: n }, () => new Array(n).fill(0));

    // Build the DP table
    for (let i = n - 1; i >= 0; i--) {
        dp[i][i] = 1;
        for (let j = i + 1; j < n; j++) {
            if (s[i] === s[j]) {
                dp[i][j] = (i + 1 === j) ? 2 : dp[i + 1][j - 1] + 2;
            } else {
                dp[i][j] = Math.max(dp[i + 1][j], dp[i][j - 1]);
            }
        }
    }

    let maxProd = 0;

    // Iterate through all possible split points i
    // Left part: s[0...i], Right part: s[i+1...n-1]
    for (let i = 0; i < n - 1; i++) {
        const leftLPS = dp[0][i];
        const rightLPS = dp[i + 1][n - 1];
        maxProd = Math.max(maxProd, leftLPS * rightLPS);
    }

    return maxProd;
}

// Example Usage:
// console.log(maxProduct("abacaba")); // Output: 9
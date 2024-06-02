package regular_expression_matching

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p

	m := len(p)
	n := len(s)

	dp := createBool2dSlice(m, n)
	dp[0][0] = true
	for i := 1; i < m; i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-2][0]
		}
		for j := 1; j < n; j++ {
			if p[i] == s[j] || p[i] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i] == '*' {
				dp[i][j] = dp[i-2][j] || ((p[i-1] == s[j] || p[i-1] == '.') && dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func createBool2dSlice(m int, n int) [][]bool {
	ans := make([][]bool, m)
	for i := range ans {
		ans[i] = make([]bool, n)
	}
	return ans
}

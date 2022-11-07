package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	dict := make(map[string]struct{})

	for _, word := range wordDict {
		dict[word] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := dict[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(dp)-1]
}

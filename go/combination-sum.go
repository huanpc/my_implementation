package main

// https://leetcode.com/problems/combination-sum

import (
	"fmt"
	"sort"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	return backtrackSol(candidates, target)
}

func backtrackSol(candidates []int, target int) [][]int {
	// outline
	// 1. create result list & combination list to store each combination made
	// 2. backtrack(startIdx, target, combination, result) updates combination which sum up to target with candidate index starting from startIdx
	//  - base case if target == 0, nothing to do anymore, add combination to result
	//  - if target < 0, do nothing, return
	//  - for each candidate, pick it and call backtrack(candidateIdx, target-candidate) to get new combination with chosen candidate. before choosing new candidate, remove current candidate from combination
	comb := make([]int, 0)
	ans := make([][]int, 0)
	backtrack(0, target, candidates, comb, &ans)
	return ans
}

func backtrack(startIdx int, target int, candidates []int, combination []int, result *[][]int) {
	if target == 0 {
		tmp := append([]int{}, combination...)
		*result = append(*result, tmp)
		return
	}
	if target < 0 {
		return
	}
	for i := startIdx; i < len(candidates); i++ {
		if target < candidates[i] {
			continue
		}
		// pick and recursively solve new target
		combination = append(combination, candidates[i])
		backtrack(i, target-candidates[i], candidates, combination, result)
		// unpick
		combination = combination[:len(combination)-1]
	}
}

// DP, bottom up solution, build from 0 to target
// dp[target] = sum(dp[target-i]), i = [1,target]
// dp[0] = 1
// need to sort to dedup, as we don't narrow candidate scope after each iteration
func dpSol(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}}
	for i := 1; i <= target; i++ {
		combine := make(map[string]bool)
		for _, cdd := range candidates {
			// i is smaller than candidate, can't form combination sum
			if i-cdd < 0 {
				continue
			}
			// pick candidate and find combination of i-cdd to form new combination of i
			for _, comb := range dp[i-cdd] {
				newComb := append([]int{}, comb...)
				newComb = append(newComb, cdd)
				if combine[toString(newComb)] {
					// duplication
					continue
				}
				combine[toString(newComb)] = true
				dp[i] = append(dp[i], newComb)
			}
		}

	}
	return dp[target]
}

func toString(l []int) string {
	sort.Ints(l)
	var b strings.Builder
	for _, v := range l {
		fmt.Fprintf(&b, "%d-", v)
	}
	return b.String()
}

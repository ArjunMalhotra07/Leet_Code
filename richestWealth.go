//https://leetcode.com/problems/richest-customer-wealth/

package main

func maximumWealth(accounts [][]int) int {

	ans := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum = sum + accounts[i][j]
		}
		if sum > ans {
			ans = sum
		}

	}
	return ans

}

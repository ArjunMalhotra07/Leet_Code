//https://leetcode.com/problems/richest-customer-wealth/

package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	f := fmt.Println

	f("Richest Wealth LeetCode Problem")

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

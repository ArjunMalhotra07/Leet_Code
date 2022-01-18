//https://leetcode.com/problems/jewels-and-stones/
package main

import (
	"fmt"
	"strings"
)

func main() {
	ans := numJewelsInStones("aA", "aAAbbbb")
	fmt.Println(ans)
}

func numJewelsInStones(jewels string, stones string) int {
	sum := 0
	map1 := mappingStones(jewels)
	map2 := mappingStones(stones)

	for key, _ := range map1 {
		sum += (map2[key])
	}
	return (sum)

}

func mappingStones(test string) map[string]int {
	map1 := make(map[string]int)
	byteTest := []byte(test)

	for i := 0; i < len(test); i++ {
		tempS := string(byteTest[i])
		tempI := strings.Count(test, tempS)
		map1[tempS] = tempI
	}
	return map1
}

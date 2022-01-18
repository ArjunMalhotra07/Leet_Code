//https://leetcode.com/problems/first-unique-character-in-a-string/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("glo")
	check2("sunaina")
	//fmt.Println(ans)
}

func check2(str string) {
	p := fmt.Println
	map1 := make(map[string]int)
	for i := 0; i < len(str); i++ {
		tempS := string(str[i])
		tempI := strings.Count(str, tempS)

		map1[tempS] = tempI
	}
	p(map1)

	for i := 0; i < len(str); i++ {
		if map1[string(str[i])] == 1 {
			p(string(str[i]))
			break
		}
	}

}

//https://leetcode.com/problems/capitalize-the-title/

package main

import (
	"fmt"
	"strings"
)

func main() {

	f1 := fmt.Println
	title := "capiTalIze tHe titLe"
	s := strings.ToLower(title)
	words := strings.Fields(s)
	var sb strings.Builder

	for _, ans := range words {
		if len(ans) == 1 || len(ans) == 2 {
			sb.WriteString(strings.ToLower(ans))
			sb.WriteString(" ")
		} else {
			sb.WriteString(strings.Title(ans))
			sb.WriteString(" ")
		}

	}
	//f1(sb.String())
	res := sb.String()
	f1(strings.Trim(res, " "))
}

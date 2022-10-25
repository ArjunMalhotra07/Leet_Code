//https://leetcode.com/problems/capitalize-the-title/

package Strings

import (
	"fmt"
	"strings"
)

func CapitalizeTitle() {
	f := fmt.Println

	f("PROGRAM 1 : To capitalize a title")
	title := "capiTalIze tHe titLe"
	f("Orignal title :  ", title)
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
	// f(sb.String())    // Ismei space ati hai at the end jiski vajah se hume ise trim krna hai
	result := sb.String()
	f("Ans -- ", strings.TrimSpace(result))
	f()
	f()
}

// https://leetcode.com/problems/implement-strstr/

package Strings

import "fmt"

func StrstrHelper() {
	f := fmt.Println
	var haystack, needle string
	f("Enter 2 Strings ")
	f("Haystack -- ")
	fmt.Scan(&haystack)
	f("Needle -- ")
	fmt.Scan(&needle)
	ans := strStr(haystack, needle)
	if ans != "" {
		f(ans, " Exists in ", haystack)
	} else {
		f(needle, " Does not exist in ", haystack)
	}
	ans1 := str(haystack, needle)
	f(ans1, " -- Index is at which the substring starts")
}

func strStr(haystack string, needle string) string {
	needleByte := []byte(needle)
	haystackByte := []byte(haystack)
	ansArray := []byte{}

	lenNeedle := len(needle)
	lenHaystack := len(haystack)

	for i := 0; i < lenHaystack; i++ {
		if haystackByte[i] == needleByte[0] {
			// fmt.Println(i)
			ansArray = haystackByte[i : i+lenNeedle]
			tempString := string(ansArray)
			if tempString != needle {
				continue
			} else {
				return string(ansArray)
			}
		}
	}
	return ""
}

func str(haystack string, needle string) int {
	needleByte := []byte(needle)
	haystackByte := []byte(haystack)
	ansArray := []byte{}

	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}

	for i := 0; i < lenHaystack; i++ {
		if haystackByte[i] == needleByte[0] {
			// fmt.Println(i)
			ansArray = haystackByte[i : i+lenNeedle]
			tempString := string(ansArray)
			if tempString != needle {
				continue
			} else {
				return i
			}
		}
	}
	return -1
}

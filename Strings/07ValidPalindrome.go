package Strings

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func CheckAnsHelper() {
	p := fmt.Println
	p(checkAns("A man, a plan, a canal: Panama"))
}

func checkAns(test string) bool {
	p := fmt.Println
	testString := strings.ToLower(test)

	//reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	//gives and err which we dont need
	//OR
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(testString, "")
	p(processedString)
	ans := booleanAns(processedString)
	return ans
}

func booleanAns(processedString string) bool {
	byte_String := []byte(processedString)
	var temp byte

	for i, j := 0, len(byte_String)-1; i < j; i, j = i+1, j-1 {
		temp = byte_String[i]
		byte_String[i] = byte_String[j]
		byte_String[j] = temp
	}

	if string(byte_String) == processedString {
		return true

	} else {
		return false
	}
}

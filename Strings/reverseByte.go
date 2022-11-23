package Strings

import "fmt"

func Reverse() {
	f := fmt.Println
	f("PROGRAM 13 : Reverse words ")

	a := []byte{97, 99, 115, 120, 80}

	f("Reverse of Hello  ---  ", reverse("Hello"))
	f("Reverse of", string(a), "  ---  ", reverse(string(a)))
	f()
	f()

}

func reverse(teststring string) string {

	ansString := []byte(teststring)
	var temp byte

	for i, j := 0, len(ansString)-1; i < j; i, j = i+1, j-1 {
		temp = ansString[i]
		ansString[i] = ansString[j]
		ansString[j] = temp
	}
	return string(ansString)
}

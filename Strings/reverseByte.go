package main

import "fmt"

func main() {
	reverse("Hello")
	a := []byte{97, 99, 115, 120, 80}
	fmt.Println(string(a))
}

func reverse(teststring string) {

	ansString := []byte(teststring)
	var temp byte

	for i, j := 0, len(ansString)-1; i < j; i, j = i+1, j-1 {
		temp = ansString[i]
		ansString[i] = ansString[j]
		ansString[j] = temp
	}
	fmt.Println(string(ansString))
}

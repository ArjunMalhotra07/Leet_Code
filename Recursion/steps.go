package main

import "fmt"

func main() {

	fmt.Println(count(14, 0))
}

func count(n int, steps int) int {

	if n == 0 {
		return steps
	}
	if n%2 == 0 {
		return count(n/2, steps+1)
	}
	return count(n-1, steps+1)

}

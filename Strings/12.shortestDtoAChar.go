// https://leetcode.com/problems/shortest-distance-to-a-character/submissions/

package Strings

import "fmt"

func MapWords() {
	f := fmt.Println
	countMap := mapWords("aaab", 98)
	f(countMap)
	f(minLength("aaab", countMap))

	newMap := mapWords("abaa", 98)
	f(newMap)
	f(minLength("abaa", newMap))
}

func mapWords(s string, u byte) map[string][]int {
	countMap := make(map[string][]int)
	testString := []byte(s)
	length := len(s)
	tempArray := []int{}
	for i := 0; i < length-1; i++ {

		if _, found := countMap[string(testString[i])]; !found {
			tempArray = append(tempArray, i)
			for j := i + 1; j < length; j++ {
				tempStringI := string(testString[i])
				tempStringJ := string(testString[j])
				if tempStringI == tempStringJ {
					tempArray = append(tempArray, j)
				}
			}
			countMap[string(testString[i])] = tempArray[:]
		}
		tempArray = []int{}

	}
	if _, found := countMap[string(testString[length-1])]; !found {
		tempArray = append(tempArray, length-1)
		countMap[string(testString[length-1])] = tempArray
	}

	newMap := make(map[string][]int)
	newMap[string(u)] = countMap[string(u)]
	// fmt.Println(newMap)
	return newMap
}

func minLength(testString string, countMap map[string][]int) []int {
	ansArray := []int{}
	testSByte := []byte(testString)
	length := len(testSByte)

	tempArray := []int{}
	for i := 0; i < length; i++ {
		for _, arrayIn := range countMap {
			for j := 0; j < len(arrayIn); j++ {
				temp := arrayIn[j] - i
				if temp < 0 {
					temp *= -1
				}
				tempArray = append(tempArray, temp)

			}
			minDist := min(tempArray)
			ansArray = append(ansArray, minDist)
			tempArray = []int{}
		}
	}
	return ansArray
}

func min(tempArray []int) int {
	minDist := tempArray[0]
	for _, value := range tempArray {
		if value < minDist {
			minDist = value
		}
	}
	return minDist
}

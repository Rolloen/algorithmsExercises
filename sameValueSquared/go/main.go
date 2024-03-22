package main

import (
	"log"
)

func main() {

	log.Println("TTessting")
}

func same(normalArray []int, squaredArray []int) bool {
	if len(normalArray) == 0 || len(squaredArray) == 0 || len(squaredArray) < len(normalArray) {
		return false
	}
	// create the result bool as true
	result := true
	// create a map frequencyNormalMap of [int]int ([normalValue]nbOfFrequency)
	// create a map frequencySquaredMap of [int]int ([squaredValue]nbOfFrequency)
	frequencyNormalMap := make(map[int]int)
	frequencySquaredMap := make(map[int]int)
	// loop over the 1st array (normalArray)
	for _, normalValue := range normalArray {
		// add one to the frequence map
		frequencyNormalMap[normalValue]++
	}
	// loop over the 2nd array (squaredArray)
	for _, squaredValue := range squaredArray {
		// add one to the frequence map
		frequencySquaredMap[squaredValue]++
	}
	// loop over the frequenceArray
	for key, freqVal := range frequencyNormalMap {
		if frequencySquaredMap[key*key] != freqVal {
			return false
		}
	}
	return result
}

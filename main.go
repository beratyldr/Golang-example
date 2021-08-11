package main

import (
	"fmt"
)

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {
	duplicate := []string{"orange", "banana", "pie", "pie", "pie", "apple", "banana", "pie", "red", "banana", "red", "blue", "blue", "gray", "gray", "pie", "pie", "fruit", "banana", "banana", "purple", "orange"}
	count := 0
	maxDup := ""
	dup_map := dup_count(duplicate)
	for k, v := range dup_map {
		if count < v {
			maxDup = k
			count = v
		}

	}
	fmt.Printf("Item : %s , Count : %d\n", maxDup, count)

}

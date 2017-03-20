package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort map by key
	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	for key, val := range hash {
		fmt.Println(key, "->", val)
	}

	fmt.Println("-----------------------------")

	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i := range keys {
		fmt.Println(keys[i], "->", hash[keys[i]])
	}
}

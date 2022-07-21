package main

import (
	"strings"
	"wc"
)

func WordCount(s string) map[string]int {
	s2 := strings.Split(s, " ")
	var myMap map[string]int
	myMap = make(map[string]int)

	for _, v := range s2 {
		myMap[v] += 1
	}

	//return map[string]int{"x": 1}
	return myMap
}

func main() {
	wc.Test(WordCount)
}

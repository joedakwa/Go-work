package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Printf("dkhkjhd")
	fmt.Scanln(&text)
	textlow := strings.ToLower(text)
	rule1 := strings.Index(textlow, "i") == 0
	rule2 := strings.Contains(textlow, "a")
	rule3 := strings.LastIndex(textlow, "n") == len(textlow)-1
	if rule1 && rule2 && rule3 {
		fmt.Printf("Found")
	} else {
		fmt.Printf("Not Found")
	}
}

// press enter twice and the compiler will execute.

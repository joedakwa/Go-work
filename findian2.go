package main

import (
	"fmt"
	"strings"
)

func main() {

	var mystring string

	fmt.Printf("Iannnn")
	fmt.Scan(&mystring)

	// Convert to lower case to make the search case insensitive
	lower_string := strings.ToLower(mystring)
	trimmed_lower_string := strings.TrimSpace(lower_string)

	if strings.HasPrefix(trimmed_lower_string, "i") && strings.HasSuffix(trimmed_lower_string, "n") && strings.Contains(trimmed_lower_string, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}

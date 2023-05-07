package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	fmt.Println("iasdas ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	str = strings.ToLower(str)
	if strings.Contains(str, "a") && str[0] == 'i' && str[len(str)-1] == 'n' {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not found!")
	}
}

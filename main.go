package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		value := []rune(os.Args[1])
		format := os.Args[2]
		count := 0
		for _, r := range format {
			c := string(r)
			if c == "x" && count < len(value) {
				fmt.Print(string(value[count]))
				count++
			} else {
				fmt.Print(c)
			}
		}
		fmt.Println()
	}
}

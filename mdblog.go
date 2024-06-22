package main

// always main package for executable program

import (
	"fmt"
	"os"
)


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: main <input.md>")
		os.Exit(1)
	}

	//WriteHTMLFile(os.Args[1])

	entries, err := os.ReadDir("./")
    if err != nil {
		os.Exit(1)
    }
 
    for _, e := range entries {
            fmt.Println(e.Name())
    }
}

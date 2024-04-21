package main

// always main package for executable program

import (
	"fmt"
	"os"

	"mdcreator/html" // to access the functions use html.<function_name>
)


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: main <input.md>")
		os.Exit(1)
	}

	html.WriteHTMLFile(os.Args[1])
}

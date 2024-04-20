package main

import (
    "fmt"
    "os"

    "mdcreator"
)


func main() {

    if len(os.Args) < 2 {
	fmt.Println("Usage: main <input.md>")
	os.Exit(1)
    }

    mdcreator.writeHTMLFile(os.Args[1])
}

package main

// always main package for executable program

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func findMarkdown(dir string) ([]string, error) {
	var templates []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			templates = append(templates, path)
		}
		return nil
	})

	if len(templates) == 0 {
		return nil, fmt.Errorf("no markdown files found in %s", dir)
	}
	return templates, err
}

func main() {
	// Define flags
	base := flag.String("base", "", "the base template file")
	blogDir := flag.String("blog-dir", "", "the directory containing blog templates")

	// Parse the flags
	flag.Parse()

	// Use the flag values
	if *base == "" || *blogDir == "" {
		fmt.Println("Usage: app --base base --blog-dir dir/")
		return
	}

	md_files, err := findMarkdown(*blogDir)
	if err != nil {
		log.Fatalf("Error finding templates: %v", err)
	}
	for _, tmpl := range md_files {
		fmt.Println(tmpl)
	}
}

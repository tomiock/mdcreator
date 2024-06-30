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
		log.Fatalf("Error finding markdown files: %v", err)
	}

	views := "views"
	base_dir := filepath.Dir(*base)
	subfolderPath := filepath.Join(base_dir, views)

	info, err := os.Stat(subfolderPath)
	if os.IsNotExist(err) {
		os.Mkdir(subfolderPath, 0755)
		// 0755 Commonly used on web servers. The owner can read, write, execute.
		// Everyone else can read and execute but not modify the file.
	} else if err != nil {
		log.Panic("An error occurred: %v\n", err)
	} else if !info.IsDir() {
		log.Fatal("%s exists but is not a directory\n", views)
	} else {
	}

	for _, tmpl := range md_files {
		WriteHTMLFile(tmpl)
	}
}

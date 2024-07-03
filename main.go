package main

// always main package for executable program

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func ToSnakeCase(input string) string {
    re := regexp.MustCompile(`[^\w]+`)
    snake := re.ReplaceAllString(strings.ToLower(input), "_")
    return strings.Trim(snake, "_")
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

	views := "views_blog"
	base_dir := filepath.Dir(*base)
	subfolderPath := filepath.Join(base_dir, views)

	info, err := os.Stat(subfolderPath)
	if os.IsNotExist(err) {
		os.Mkdir(subfolderPath, 0755)
		// 0755 Commonly used on web servers. The owner can read, write, execute.
		// Everyone else can read and execute but not modify the file.
	} else if err != nil {
		log.Panicf("An error occurred: %v\n", err)
	} else if !info.IsDir() {
		log.Fatalf("%s exists but is not a directory\n", views)
	} else { 
	}

	for _, tmpl := range md_files {
		title_post := TransformFileName(tmpl[len(*blogDir)+1:len(tmpl)-3])
		_title_post := ToSnakeCase(title_post)
		post_path := filepath.Join(subfolderPath, title_post + ".tmpl")
		WriteTmplFile(tmpl, post_path)
		fmt.Println(post_path)
		fmt.Println(_title_post)
		fmt.Println(title_post)
	}
}

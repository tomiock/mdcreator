package main

// always main package for executable program

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

func generateGetRouteCode(post_title, route_name, base_path, file_path string) string {
	return fmt.Sprintf(`e.GET("/%s", func(c echo.Context) error {
	t, err := template.ParseFiles("../%s", "../%s")
	if err != nil {
		panic(err)
	}

	res := map[string]interface{}{
		"Title": "%s",
	}
	return t.Execute(c.Response().Writer, res)
})`, route_name, base_path, file_path, post_title)
}

func updateMainFile(file_path string, route string) error {
	regex, err := regexp.Compile(`^\w+\.Logger\.Fatal\(\w+\.Start\(":\d+"\)\)`)
	if err != nil {
		return err
	}

	var last_lines []string
	var lines []string
	matched := false

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("Error opening main file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if regex.MatchString(line) {
			matched = true
		}
		if matched {
			// store the line in a buffer
			last_lines = append(last_lines, line)
		} else {
			lines = append(lines, line)
		}

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	file, err = os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	_, err = writer.WriteString(route + "\n\n")
	if err != nil {
		return err
	}

	for _, line := range last_lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}	
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func main() {
	// Define flags
	base := flag.String("base", "", "the base template file")
	blogDir := flag.String("blog-dir", "", "the directory containing blog templates")
	mainFile := flag.String("main", "", "the main file to write the routes to")

	// Parse the flags
	flag.Parse()

	// Use the flag values
	if *base == "" || *blogDir == "" || *mainFile == "" {
		fmt.Println("Usage: app --base base --blog-dir dir/ --main main.go")
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
		title_post := tmpl[len(*blogDir)+1 : len(tmpl)-3]
		_title_post := FileName_to_snake_case(title_post)
		post_path := filepath.Join(subfolderPath, _title_post+".tmpl")

		WriteTmplFile(tmpl, post_path)

		route_str := generateGetRouteCode(title_post, _title_post, *base, post_path)

		err = updateMainFile(*mainFile, route_str)
		if err != nil {
			log.Fatalf("Error updating main file: %v", err)
		}
	}
}


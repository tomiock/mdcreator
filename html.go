package main

import (
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md string) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	md_file := []byte(md)
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md_file)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func TransformFileName(fileName string) string {
	list_string := strings.Fields(strings.ToLower(fileName))
	return strings.Join(list_string, "_")
}

func WriteHTMLFile(args string) {

	input, err := os.ReadFile(args)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	file_name := TransformFileName(args[:len(args)-3])

	file, err := os.Create(file_name + ".html")
	if err != nil {
		log.Panicf("Error creating file: %v", err)
	}

	html := mdToHTML(string(input))

	header := []byte("{{ define \"content\"}}\n")

	_, err = file.Write(header)
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}

	_, err = file.Write(html)
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}

	_, err = file.Write([]byte("{{ end }}"))
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}
}

func WriteTmplFile(args string, file_path string) {
	input, err := os.ReadFile(args)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	// if the file already exists, it overwrites it
	file, err := os.Create(file_path)
	if err != nil {
		log.Panicf("Error creating file: %v", err)
	}

	_, err = file.Write([]byte("{{ define \"content\"}}\n"))
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}

	html := mdToHTML(string(input))
	_, err = file.Write(html)
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}

	_, err = file.Write([]byte("{{ template \"content\" . }}"))
	if err != nil {
		log.Panicf("Error writing to file: %v", err)
	}
}

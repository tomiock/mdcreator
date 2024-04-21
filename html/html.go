package html

import (
	"fmt"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}


func TranformFileName(fileName string) string {
	list_string := strings.Fields(fileName)

	for idx := range list_string {
		caser := cases.Title(language.English)
		list_string[idx] = caser.String(list_string[idx])
	}

	return strings.Join(list_string, "")
}

func WriteHTMLFile(args string) {

    input, err := os.ReadFile(args)
    if err != nil {
	fmt.Println("Error reading file:", err)
	panic(err)
    }

    file_name := TranformFileName(os.Args[1][:len(args)-3])

    file, err := os.Create(file_name + ".html")
    if err != nil {
	fmt.Println("Error creating file:", err)
	panic(err)
    }

    html := mdToHTML(input)

    header := []byte("{{ block " + file_name + " . }}\n<!DOCTYPE html>\n")
    _, err = file.Write(header)
    _, err = file.Write(html)
    _, err = file.Write([]byte("{{ end }}"))
}

package tests

import (
	"bufio"
	"mdcreator/html"
	"os"

	"testing"
)


func TestMdcreator(t *testing.T) {
    f, err := os.Create("test.md")
    if err != nil {
        t.Fatal(err)
    }

    data := []byte("# Title\n## Second Title\nTesting the main function")
    _, err = f.Write(data)

    if err != nil {
        t.Fatal(err)
    }

    html.WriteHTMLFile("test.md")

    file_read, err := os.Open("test.html")
    if err != nil {
        t.Fatal(err)
    }

    read_lines := []string{"{{ block \"test\" . }}",
    "<!DOCTYPE html>",
    "<h1 id=\"title\">Title</h1>", 
    "",
    "<h2 id=\"second-title\">Second Title</h2>", 
    "",
    "<p>Testing the main function</p>", 
    "{{ end }}"}

    scanner := bufio.NewScanner(file_read)
    idx := 0
    for scanner.Scan() {
        if scanner.Text() != read_lines[idx] {
            t.Fatalf("Expected %s, got %s", read_lines[idx], scanner.Text())
        }
        idx++
    }

    if err := scanner.Err(); err != nil {
        t.Fatal(err)
    }
}

func TestTranformFileName(t *testing.T) {
    expected := "test_filename_with_a_complex_name.md"

    result := html.TransformFileName("Test FileName with a Complex Name.md")

    if result != expected {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

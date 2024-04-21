package tests

import (
	"bufio"
	"mdcreator/html"
	"os"

	"testing"
)


func TestMdcreator(t *testing.T) {
    test_file := "test.md"

    f, err := os.Create(test_file)
    if err != nil {
        t.Fatal(err)
    }

    data := []byte("# Title\n## Second Title\nTesting the main function")
    _, err = f.Write(data)

    if err != nil {
        t.Fatal(err)
    }

    html.WriteHTMLFile(test_file)

    file_read, err := os.Open("Test.html")
    if err != nil {
        t.Fatal(err)
    }

    read_lines := []string{"{{ block Test . }}",
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
    expected := "TestFilenameWithAComplexName.md"

    result := html.TranformFileName("Test FileName with a Complex Name.md")

    if result != expected {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

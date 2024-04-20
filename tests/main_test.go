package tests

import (
    "os"
    "mdcreator"

    "testing"
)

func TestMain(t *testing.T) {
    file := "test.md"  

    f, err := os.Create(file)
    if err != nil {
        t.Fatal(err)
    }

    data := []byte("# Title\n## Second Title\nTesting the main function")
    _, err = f.Write(data)

    if err != nil {
        t.Fatal(err)
    }
    mdcreator.writeHTMLFile(file)
}

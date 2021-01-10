package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
)

func main() {
	file := "test.md"

	content, readErr := ioutil.ReadFile(file)

	if readErr != nil {
		log.Fatalf("%s file not found", file)

	}

	html := markdown.ToHTML(content, nil, nil)
	fmt.Printf(string(html))

	fileOut := "test.html"
	writeErr := ioutil.WriteFile(fileOut, html, 0644)

	if writeErr != nil {
		log.Fatalf("Could not write to %s", fileOut)

	}

	fmt.Printf("HTML outputted to %s \n", fileOut)
}

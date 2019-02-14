package main

import (
	"fmt"
	"io"
	"os"

	"github.com/syumai/printpdf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give a URL to print to PDF.")
		os.Exit(1)
	}

	pdfReader, err := printpdf.NewReader(os.Args[1])
	if err != nil {
		fmt.Printf("unexpected error: %v", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, pdfReader)
	if err != nil {
		fmt.Printf("unexpected error: %v", err)
		os.Exit(1)
	}
}

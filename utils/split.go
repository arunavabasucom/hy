package utils

import (
	"fmt"
	"os"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func Split() {
    // Input and output file paths
    inputFile := "merged.pdf"
    outputDir := "output"

    // Create the output directory if it doesn't exist
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        fmt.Printf("Error creating output directory: %v\n", err)
        return
    }
     
    // Split the PDF into individual pages
    if err := pdfcpu.SplitFile(inputFile, outputDir, 1, nil); err != nil {
        fmt.Printf("Error splitting PDF: %v\n", err)
        return
    }

    fmt.Println("PDF split successfully.")
}


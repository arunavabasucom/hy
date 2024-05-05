package main

import (
	"fmt"
	"os"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func mergePDFs(outputPath string, inputPaths ...string) error {
	
	if len(inputPaths) < 2 {
		return fmt.Errorf("must provide at least two input PDFs to merge")
	}

	// Merge input PDFs
	if err := pdfcpu.MergeCreateFile(inputPaths, outputPath, false, nil); err != nil {
		return fmt.Errorf("error merging PDFs: %v", err)
	}

	return nil
}

func main() {
	outputPath := "merged.pdf"
	inputPaths := []string{"input1.pdf", "input2.pdf"}

	if err := mergePDFs(outputPath, inputPaths...); err != nil {
		fmt.Fprintf(os.Stderr, "Error merging PDFs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PDFs merged successfully.")
}

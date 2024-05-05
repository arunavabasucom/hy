package utils

import (
	"fmt"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePDFs(outputPath string, inputPaths ...string) error {
	/*
	 merge pdfs 
	*/
	if len(inputPaths) < 2 {
		return fmt.Errorf("must provide at least two input PDFs to merge")
	}

	// Merge input PDFs
	if err := pdfcpu.MergeCreateFile(inputPaths, outputPath, false, nil); err != nil {
		return fmt.Errorf("error merging PDFs: %v", err)
	}

	return nil
}


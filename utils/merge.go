package utils

import (
	"fmt"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

/*
Merge two pdfs
*/
func MergePDFs(outputPath string, inputPaths ...string) error {
	if len(inputPaths) < 2 {
		return fmt.Errorf("must provide at least two input PDFs to merge")
	}
	if err := pdfcpu.MergeCreateFile(inputPaths, outputPath, false, nil); err != nil {
		return fmt.Errorf("error merging PDFs: %v", err)
	}

	return nil
}

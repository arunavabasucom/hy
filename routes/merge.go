package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
    "hy/utils"
	"github.com/gin-gonic/gin"
)

func MergePDFHandler(c *gin.Context) {
	// Get the uploaded files
	file1, err := c.FormFile("file1")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file1 not found: %v", err))
		return
	}
	file2, err := c.FormFile("file2")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file2 not found: %v", err))
		return
	}

	// Save the uploaded files
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create upload directory: %v", err))
		return
	}
	filePath1 := filepath.Join(uploadDir, file1.Filename)
	filePath2 := filepath.Join(uploadDir, file2.Filename)
	fmt.Println(filePath1)
	fmt.Println(filePath2)
	if err := c.SaveUploadedFile(file1, filePath1); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to save file1: %v", err))
		return
	}
	if err := c.SaveUploadedFile(file2, filePath2); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to save file2: %v", err))
		return
	}

	// Merge the uploaded PDF files
	outputPath := filepath.Join(uploadDir, "merged.pdf")
	if err := utils.MergePDFs(outputPath, filePath1, filePath2); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to merge PDFs: %v", err))
		return
	}

	// Send the merged PDF file to the client
	c.File(outputPath)
}

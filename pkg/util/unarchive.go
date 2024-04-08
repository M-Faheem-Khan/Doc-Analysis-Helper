package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UnArchive(fname string, dst string) {
	fmt.Printf("Unzipping %s.\n", fname)

	archive, err := zip.OpenReader(fname)
	if err != nil {
		fmt.Printf("Error opening %s\n %s\n", fname, err)
		os.Exit(1)
	}

	defer archive.Close()

	// Extract the files from the zip
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		// Check if the file is a directory
		if f.FileInfo().IsDir() {
			// Create the directory
			fmt.Printf("creating directory %s\n", f.Name)
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				fmt.Printf("Error creating directory %s\n %s\n", f.Name, err)
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s\n %s\n", f.Name, err)
		}

		// Create an empty destination file
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			panic(err)
		}

		// Close the files
		dstFile.Close()
		srcFile.Close()

	}

	fmt.Printf("Extracted %s to %s.\n", fname, dst)
}

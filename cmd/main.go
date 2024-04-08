package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/iocs"
	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/util"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Error: No file given for analysis.\n")
		fmt.Printf("Usage: %s <file-to-analyze>", os.Args[0])
		os.Exit(1)
	}

	fname := os.Args[1]

	fmt.Printf("Starting analysis on %s\n", fname)

	dst := "analysis_extract"

	util.UnArchive(fname, dst)

	// Iterate over files and look for External Connections
	err := filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// Check for IOCs
			iocs.OutboundConnections(content, path)

		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning %s\n %s\n", fname, err.Error())
	}

	// outbound connections (http, ftp, share, ip)
	// vba macros
	// masked files (theme to make the user disable security mode)

	// Given a file
	// $ program excel.xlsx
	// unzips the excel.xlsx
	// looks for iocs
	// generate report

}

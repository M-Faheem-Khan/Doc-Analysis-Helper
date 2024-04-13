package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/iocs"
	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/report"
	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/util"
)

func getFileHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Error: No file given for analysis.\n")
		fmt.Printf("Usage: %s <file-to-analyze>", os.Args[0])
		os.Exit(1)
	}

	fname := os.Args[1]

	var r report.REPORT

	r.FileName = fname
	r.AnalysisDate = time.Now().String()
	r.SHA256Hash = getFileHash(fname)

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
			ioc := iocs.OutboundConnections(content, path)
			if ioc.IOC != "" {
				r.IndicatorsOfCompromise.OutBoundConnections = append(r.IndicatorsOfCompromise.OutBoundConnections, ioc)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning %s\n %s\n", fname, err.Error())
	}

	r.WriteReport("report.md")

	fmt.Printf("Analysis Complete")

	// outbound connections (http, ftp, share, ip)
	// vba macros
	// masked files (theme to make the user disable security mode)

	// Given a file
	// $ program excel.xlsx
	// unzips the excel.xlsx
	// looks for iocs
	// generate report

}

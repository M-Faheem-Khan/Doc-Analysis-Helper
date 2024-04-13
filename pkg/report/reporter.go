package report

import (
	"fmt"
	"os"

	"github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/iocs"
)

type IOCS struct {
	OutBoundConnections []iocs.OutBoundConnection
}

type REPORT struct {
	FileName               string
	SHA256Hash             string
	AnalysisDate           string
	IndicatorsOfCompromise IOCS
}

func (r *REPORT) PrintReport() {
	fmt.Printf("File: %s\n", r.FileName)
	fmt.Printf("SHA256 Hash: %x\n", r.SHA256Hash)
	fmt.Printf("Date Analyzed: %s\n", r.AnalysisDate)

	fmt.Printf("\n\n\n")
	for _, obc := range r.IndicatorsOfCompromise.OutBoundConnections {
		fmt.Printf("IOC: %s\n", obc.IOC)
		fmt.Printf("File Path: %s\n", obc.IOC)
	}
}

func (r *REPORT) WriteReport(fname string) {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}

	f.Write([]byte("**File:** " + r.FileName + "  \n"))
	f.Write([]byte("**SHA256 Hash:** " + r.SHA256Hash + "  \n"))
	f.Write([]byte("**Date Analyzed:** " + r.AnalysisDate + "  \n"))

	f.Write([]byte("\n\n---------------\n\n\n"))

	f.Write([]byte("## Indicators of Compromise  \n\n"))
	f.Write([]byte("### Out Bound Connections  \n\n"))
	for _, obc := range r.IndicatorsOfCompromise.OutBoundConnections {
		f.Write([]byte("**IOC:** " + obc.IOC + "  \n"))
		f.Write([]byte("**File Path:** `" + obc.Path + "`  \n\n"))
	}

	f.Write([]byte("\n"))
}

package report

import "github.com/m-faheem-khan/Doc-Analysis-Helper/pkg/iocs"

type IOCS struct {
	OutBoundConnections []iocs.OutBoundConnection
}

func GenerateReport() {

	// File: xxx
	// SHA256 Hash: xxx
	// Date Analyze: xxx

	// ----- IOCs -----
	// ## OutBoundConnections
	// Potential x outbound connections were found.
	// PATH
	// IOC
	// ...
	// PATH
	// IOC
	// ----- IOCs -----

}

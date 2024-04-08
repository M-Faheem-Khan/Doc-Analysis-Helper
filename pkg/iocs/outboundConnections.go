package iocs

import (
	"fmt"
	"regexp"
	"strings"
)

type OutBoundConnection struct {
	IOC  string
	Path string
}

func OutboundConnections(content []byte, path string) OutBoundConnection {
	var obc OutBoundConnection

	pattern := regexp.MustCompile("TargetMode=\"External")
	if pattern.Match([]byte(content)) {
		fmt.Printf("OutBound Connection - Match: %s\n", path)
		// Print the url/ip that is the match
		pattern = regexp.MustCompile("Target=\"(http|file|).*\"")
		match := string(pattern.Find(content))
		matches := strings.Split(match, "><")
		for _, match := range matches {

			// External Mapping only
			if strings.Contains(match, "External") {
				// Clean up
				match = strings.ReplaceAll(match, "TargetMode=\"External\"", "")
				match = string(pattern.Find([]byte(match)))
				match = strings.ReplaceAll(match, "Target=\"", "")
				match = strings.TrimSuffix(match, "\"")

				obc.Path = path
				obc.IOC = match
				return obc
			}
		}
	}
	return obc
}

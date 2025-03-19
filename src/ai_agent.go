package main

import "strings"

func analyzeError(logs string) string {
	if strings.Contains(logs, "database") {
		return "Possible database connection failure."
	} else if strings.Contains(logs, "timeout") {
		return "Network timeout detected."
	}
	return "Unknown error. Further analysis required."
}

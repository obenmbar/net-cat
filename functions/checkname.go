package netino

import "strings"

// Add Checkname function to validate client name format and length
func Checkname(name string) bool {
	isvalide := true
	if name == "" || strings.TrimSpace(name) == "" || strings.Trim(name, "\t") == "" || len(name) > 30 {
		isvalide = false
	}
	for _, i := range name {
		if i < 32 || i > 126 {
			isvalide = false
			break
		}
	}

	return isvalide
}

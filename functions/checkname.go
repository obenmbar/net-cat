package netino

import "strings"

func Check(name string) bool {
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

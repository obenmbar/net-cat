package netino

import (
	"os"
)

// Add Logo function to read and return ASCII logo from file
// Fallback with error message when logo file cannot be read
func Logo() string {
	netlogo, err := os.ReadFile("netcat.txt")
	if err != nil {
		return "ereue quant lire logo"
	}
	return string(netlogo)
}

package netino

import (
	
	"os"
)

func Logo() string {
	netlogo, err := os.ReadFile("netcat.txt")
	if err != nil {
		return "ereue quant lire logo"
	}
	return string(netlogo)
}

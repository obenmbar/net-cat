package netino

import "time"

// Refactor: translate console messages from French to English
func Timeino() string {
	res := time.Now().Format("2006-01-02 15:04:05")
	return res
}

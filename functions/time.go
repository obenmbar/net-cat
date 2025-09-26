package netino

import "time"

func Timeino() string {
	res := time.Now().Format("2006-01-02 15:04:05")
	return res
}

package netino

// Add ChecMessege function to validate printable characters in a message
func ChecMessege(msg string) bool {
	for _, i := range msg {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}

package netino

func ChecMessege(msg string) bool {
	for _, i := range msg {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true 
}

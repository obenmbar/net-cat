package netino



func Checknamedansmap ( name string , s *Server) bool {
	for v := range s.clients {
		if v == name {
			return false
		}
	}
	return true 
}

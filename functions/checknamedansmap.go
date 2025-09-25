package netino

import "net"

func Checknamedansmap ( name string , s *Server) bool {
	for v := range s.clients {
		if v == name {
			return false
		}
	}
	return true 
}
func Checkconn (con net.Conn, s *Server) bool{
	for  _, v := range s.clients {
		if v.cone == con {
			return true
		}
	}
	return false 
}
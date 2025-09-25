package netino



func Broadcast(s *Server, msg string, name string) {
s.mu.Lock()
defer s.mu.Unlock()
for na, val := range s.clients {
	if na != name {
		val.cone.Write([]byte(msg))
	}		
	}
	s.history= append(s.history, msg)
}
package netino


func Broadcast(s *Server, msg string ){
s.mu.Lock()
defer s.mu.Unlock()
for _, val := range s.clients {
	val.cone.Write([]byte(msg))
}
s.history = append(s.history, msg)
}
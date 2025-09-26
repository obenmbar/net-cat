package netino



func Addclients(s *Server, c *Client, msg string, Numbertotale *int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[c.name] = c
	*Numbertotale= *Numbertotale+1
}

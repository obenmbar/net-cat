package netino



func Addclients(s *Server, c *Client, msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[c.name] = c
}

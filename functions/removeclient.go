package netino



func Removeclients(s *Server, c *Client,msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, c.name)
	s.history= append(s.history, msg)
}
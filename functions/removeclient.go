package netino

func Removeclients(s *Server, c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
		delete(s.clients, c.name)
	
}

package netino

func Removeclients(s *Server, c *Client, numbertotale *int) {
	s.mu.Lock()
	defer s.mu.Unlock()
		delete(s.clients, c.name)
		*numbertotale -= 1
	
	
}

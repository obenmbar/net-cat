package netino

// Add Sendhistory function to send chat history to a client
func Sendhistory(s *Server, c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, va := range s.history {
		c.cone.Write([]byte(va))
	}
}

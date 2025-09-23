package netino

func Sendhistory(s *Server, c *Client) {
	s.mu.Lock()
	for _, va := range s.history {
		c.cone.Write([]byte(va))
	}
s.mu.Unlock()
}
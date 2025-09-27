package netino

// Add Addclients function to register new clients and update total count
func Addclients(s *Server, c *Client, totalClients *int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[c.name] = c
	*totalClients = *totalClients + 1
}

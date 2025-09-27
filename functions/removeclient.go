package netino

// Implement Removeclients function to safely remove a client
// Decrement total clients count with proper synchronization
func Removeclients(s *Server, c *Client, totalClients *int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, c.name)
	*totalClients -= 1
}

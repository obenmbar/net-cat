package netino

import "fmt"

func Removeclients(s *Server, c *Client,msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, c.name)
	s.history= append(s.history, fmt.Sprintf("%shas left the chat",c.name))
	Broadcast(s, msg)
}
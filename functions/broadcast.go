package netino

import (
	"fmt"
	"time"
)

// Add Broadcast function to send messages to all connected clients
// Maintain chat history by appending each broadcasted message
func Broadcast(s *Server, msg string, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for na, val := range s.clients {
		if na != name {
			val.cone.Write([]byte("\n" + msg))
			val.cone.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), na)))
		} else {
			val.cone.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), name)))
		}
	}
	s.history = append(s.history, msg)
}

package netino

import "fmt"


func Addclients(s *Server, c *Client, msg string){
	s.mu.Lock() 
	defer s.mu.Unlock()
	s.clients[c.name]= c
   s.history = append(s.history, fmt.Sprintf("%shas joind the chat",c.name))
   Sendhistory(s, c)
 

}
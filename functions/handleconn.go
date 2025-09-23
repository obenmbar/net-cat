package netino

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	cone net.Conn
	name string
}
type Server struct {
	clients map[string]*Client
	history []string
	mu      sync.Mutex
}

var GlobalServer = &Server{
	clients: make(map[string]*Client),
	history: []string{},
}

func Handleconn(con net.Conn) {
	defer con.Close()
	con.Write([]byte("Welcome to TCP-Chat!\n"))
	con.Write([]byte(Logo() + "\n"))
	con.Write([]byte("[ENTER YOUR NAME]:"))
	var line string

	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		line = scanner.Text()
		if Check(line) {
			con.Write([]byte("mrhba bik asi " + line + "\n"))

			Addclients(GlobalServer, &Client{cone: con, name: line}, line)
			Broadcast(GlobalServer, line)

		} else {
			con.Write([]byte("3awd ktb smya madokan fla tadrt arays"))
		}
		continue

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("erreur in scanner,", err)
	}
}

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
	con.Write([]byte("[ENTER YOUR NAME]: "))

	scanner := bufio.NewScanner(con)
	var client *Client

	// الخطوة الأولى: نسجل الاسم
	for scanner.Scan() {
		line := scanner.Text()

		// التحقق من الاسم
		if !Check(line) {
			con.Write([]byte("3awd ktb smya madokan fla tadrt arays \n"))
			continue
		}
		if !Checknamedansmap(line, GlobalServer) {
			con.Write([]byte("had smya deja kayna ktb smya jdida \n"))
			continue
		}

		client = &Client{cone: con, name: line}
		Sendhistory(GlobalServer, client)
		Addclients(GlobalServer, client, line)
		Broadcast(GlobalServer, fmt.Sprintf("%s has joined the chat \n", line), line)

		break
	}

	// الخطوة الثانية: المساجات ديالو
	for scanner.Scan() {
		line := scanner.Text()
		Broadcast(GlobalServer, fmt.Sprintf("%s: %s\n", client.name, line), client.name)

	}
	if con.Close() != nil {
		Removeclients(GlobalServer, client, fmt.Sprintf("%s has left the chat", client.name))
		Broadcast(globalsserver,fmt.Sprintf("%s has left the chat\n",client.name)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("erreur in scanner,", err)
	}
}

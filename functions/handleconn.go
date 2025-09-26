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

var (
	Time         = Timeino()
	Numbertotale *int
	num          int = 0
)

func Handleconn(con net.Conn) {
	defer con.Close()
	con.Write([]byte("Welcome to TCP-Chat!\n"))
	con.Write([]byte(Logo() + "\n"))
	con.Write([]byte("[ENTER YOUR NAME]: "))

	scanner := bufio.NewScanner(con)
	var client *Client

	Numbertotale = &num
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
		if num < 10 {
			Sendhistory(GlobalServer, client)
			Addclients(GlobalServer, client, line, Numbertotale)
			Broadcast(GlobalServer, fmt.Sprintf("%s has joined the chat...\n", line), line)

		} else {
			con.Write([]byte("Server full. Max connections reached"))
			con.Close()
		}
	
		break
	}

	// الخطوة الثانية: المساجات ديالو

	for scanner.Scan() {
		line := scanner.Text()
		Broadcast(GlobalServer, fmt.Sprintf("[%s][%s]:%s\n", Time, client.name, line), client.name)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("erreur in scanner,", err)
		if len(GlobalServer.clients) > 0 {
			Removeclients(GlobalServer, client, Numbertotale)
		 Broadcast(GlobalServer, fmt.Sprintf("%s has left the chat... \n", client.name), client.name)
		} else {
			fmt.Println("wahd ex khrj mn chat")
		}

	} else {
		if len(GlobalServer.clients) > 0 {
			fmt.Println("client khrj mn server")
			Removeclients(GlobalServer, client, Numbertotale)
			 Broadcast(GlobalServer, fmt.Sprintf("%s has left the chat... \n", client.name), client.name)
		} else {
			fmt.Println("wahd lmjhol khrj mn chat")
		}
	}
}

package netino

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
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
	Time        = Timeino()
	MaxClients  *int
	clientCount int = 0
	namee       string
)

// Add Handleconn function to manage client connections
// Implement name validation, history sending, and broadcasting
func Handleconn(con net.Conn) {
	defer con.Close()
	con.Write([]byte(Cyan + "Welcome to TCP-Chat!\n" + Reset))
	con.Write([]byte(Cyan + Logo() + "\n" + Reset))
	con.Write([]byte(Cyan + "[ENTER YOUR NAME]: " + Reset))

	scanner := bufio.NewScanner(con)
	var client *Client
	MaxClients = &clientCount

	for scanner.Scan() {
		namee = scanner.Text()

		if !Checkname(namee) {
			con.Write([]byte(Red + "[ENTER YOUR NAME]\n" + Reset))
			continue
		}
		if !Checknamedansmap(namee, GlobalServer) {
			con.Write([]byte(Red + "ENTER ANOTHER NAME \n" + Reset))
			continue
		}

		if clientCount < 10 {
			client = &Client{cone: con, name: namee}
			Sendhistory(GlobalServer, client)
			Addclients(GlobalServer, client, MaxClients)
			Broadcast(GlobalServer, fmt.Sprintf(Green+"%s has joined the chat...\n", namee)+Reset, namee)

		} else {
			con.Write([]byte(Red + "Server full. Max connections" + Reset))
			con.Close()
		}

		break
	}

	for scanner.Scan() {
		messege := scanner.Text()
		if ChecMessege(messege) {
			Broadcast(GlobalServer, fmt.Sprintf("[%s][%s]:%s\n", Time, client.name, messege), client.name)
		} else {
			con.Write([]byte(Red + "your are write an imprimptable cracter\n" + Reset))
			con.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.name)))
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(Red+"erreur in scanner,"+Reset, err)
		if client != nil {
			Removeclients(GlobalServer, client, MaxClients)
			Broadcast(GlobalServer, Yellow+fmt.Sprintf("%s has left the chat... \n", client.name)+Reset, client.name)
		} else {
			fmt.Println("client deconnect before joining")
		}

	} else {
		if client != nil {
			fmt.Println("client deconnect after joining")
			Removeclients(GlobalServer, client, MaxClients)
			Broadcast(GlobalServer, Yellow+fmt.Sprintf("%s has left the chat... \n", client.name)+Reset, client.name)
		} else {
			fmt.Println("client deconnect before joining")
		}
	}
}

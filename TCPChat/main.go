package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	netino "netino/functions"
)

// Add TCP server entrypoint with default port handling
// Implement connection listener and delegate client handling to netino package

var port string
func main() {
	const defaultport string = "8989"

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return

	} else if len(os.Args) == 1 {
		port = defaultport
	} else {

		_, err := strconv.Atoi(os.Args[1])
		port = os.Args[1]
		if err != nil {
			fmt.Println("this port number is not valid")
			log.Fatal(err)
		}
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("error in listener:", err)
		return
	}
	fmt.Println("Listening on the port :" + port)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error in listener accept:", err)
			continue
		}
		go netino.Handleconn(conn)

	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	netino "netino/functions"
)

var port string

func main() {
	const defaultport string = "8989"

	if len(os.Args) > 2 {

		fmt.Println("[USAGE]: ./TCPChat $port")
		return

	} else if len(os.Args) <= 1 {
		port = defaultport
	} else {

		_, err := strconv.Atoi(os.Args[1])
		port = os.Args[1]
		if err != nil {
			fmt.Println("cette port n'est vpas exist")
			log.Fatal(err)
		}
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("erreur dans listen :", err)
		return
	}
	fmt.Println("Listening on the port :" + port)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("erroro in LIstner accept :", err)
			continue
		}
		go netino.Handleconn(conn)

	}
}

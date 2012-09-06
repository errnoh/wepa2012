package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

var port int

func init() {
	rand.Seed(time.Now().UnixNano())
	port = 23000 + rand.Intn(10000)
}

func tcpserver() {
	var listener net.Listener
	var err error

	listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		println(err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var incoming string
	var err error
	var knockknock = []string{"Knock knock!", "Who's there?", "Lettuce", "Lettuce who?", "Lettuce in! it's cold out here! Bye."}

	for {
		if len(knockknock) > 0 {
			fmt.Fprintln(conn, fmt.Sprintln(knockknock[0]))
		} else {
			break
		}
		knockknock = knockknock[1:]

	fromClient:
		if incoming, err = readFromConn(conn); err != nil {
			if err.Error() != "EOF" {
				fmt.Println(err)
			}
			break
		}
		incoming = strings.TrimSpace(incoming)
		fmt.Println("Client:", incoming)
		if incoming != knockknock[0] {
			fmt.Fprintln(conn, fmt.Sprintf("You're supposed to ask \"%s\"\r\n", knockknock[0]))
			goto fromClient
		}
		knockknock = knockknock[1:]
	}
	conn.Close()
}

func readFromConn(conn net.Conn) (message string, err error) {
	message, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	message = strings.TrimSpace(message)
	return
}

func main() {
	tcpserver()
}

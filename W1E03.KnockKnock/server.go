// W1E03
//
// Knockknock server that tries to chat the following conversation:
// "Knock knock!"
// "Who's there?"
// "Lettuce"
// "Lettuce who?"
// "Lettuce in! it's cold out here! Bye."
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

// NOTE: Init ajetaan ennen muita osia ohjelman alussa.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Starts a server that listens to incoming tcp connections
func tcpserver(port int) {
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

// Handler for incoming tcp connections.
// Goes through the predefined lines and closes the connection after the last one.
func handleConnection(conn net.Conn) {
	var message string
	var err error
	var lines = []string{"Knock knock!", "Who's there?", "Lettuce", "Lettuce who?", "Lettuce in! it's cold out here! Bye."}

	for {
		// Send a line to the client
		if len(lines) > 0 {
			fmt.Fprintln(conn, fmt.Sprintln(lines[0]))
		} else {
			break
		}
		lines = lines[1:]

	fromClient:
		// Read message from the client
		if message, err = readFromConn(conn); err != nil {
			if err.Error() != "EOF" {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("Client:", message)
		if message != lines[0] {
			fmt.Fprintln(conn, fmt.Sprintf("You're supposed to ask \"%s\"\r\n", lines[0]))
			goto fromClient
		}
		lines = lines[1:]
	}
	conn.Close()
}

// Read a message from connection or timeout after 500ms
func readFromConn(conn net.Conn) (string, error) {
	c := make(chan *message)
	go func(c chan *message) {
		m := new(message)
		m.text, m.err = bufio.NewReader(conn).ReadString('\n')
		m.text = strings.TrimSpace(m.text)
		c <- m
	}(c)

	var m *message
	select {
	case m = <-c:
		// Go defaults to break
	case <-time.After(500 * time.Millisecond):
		return "", errors.New("Connection timed out")
	}
	return m.text, m.err
}

// NOTE: Channels can only return one variable so we need to use a struct to hold multiple items
type message struct {
	text string
	err  error
}

func main() {
	port := 23000 + rand.Intn(10000)
	fmt.Println("Running knockknock server in port", port)
	tcpserver(port)
}

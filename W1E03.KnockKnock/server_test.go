package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
)

var port int

func init() {
	port = 23000 + rand.Intn(10000)
	// Start the server
	go tcpserver(port)
}

func TestKnockKnockClientCorrect(t *testing.T) {
	var message string
	var err error

	// Create a tcp connection to the server
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := []string{"Who's there?", "Lettuce who?"}

	for {
		// Read a message from the server
		if message, err = readFromConn(conn); err != nil {
			t.Error(err)
		}
		// Print the received message to stdout
		fmt.Println("Server:", message)

		// End the connection if message ends in "Bye." or we run out of lines to say.
		if strings.HasSuffix(message, "Bye.") || len(lines) == 0 {
			break
		}
		// Reply to server with next line.
		fmt.Fprintln(conn, lines[0])
		lines = lines[1:]
	}
	fmt.Println()
	conn.Close()
}

func TestKnockKnockClientWrong(t *testing.T) {
	var message string
	var err error

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := []string{"What?", "Who's there?", "ugh...", "Lettuce who?", "..and?"}

	for {
		if message, err = readFromConn(conn); err != nil {
			t.Error(err)
		}
		message = strings.TrimSpace(message)
		fmt.Println("Server:", message)
		if strings.HasSuffix(message, "Bye.") {
			break
		}
		if len(lines) == 0 {
			break
		}
		fmt.Fprintln(conn, lines[0])
		lines = lines[1:]
	}
	fmt.Println()
	conn.Close()
}

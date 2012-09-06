package main

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestKnockKnockClientCorrect(t *testing.T) {
	var got string
	var err error

	go tcpserver()
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := []string{"Who's there?", "Lettuce who?"}

	for {
		if got, err = readFromConn(conn); err != nil {
			t.Error(err)
		}
		got = strings.TrimSpace(got)
		fmt.Println("Server:", got)
		if strings.HasSuffix(got, "Bye.") {
			break
		}
		if len(lines) == 0 {
			break
		}
		fmt.Fprintln(conn, lines[0])
		lines = lines[1:]
	}
	conn.Close()
}

func TestKnockKnockClientWrong(t *testing.T) {
	var got string
	var err error

	go tcpserver()
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := []string{"What?", "Who's there?", "ugh...", "Lettuce who?", "..and?"}

	for {
		if got, err = readFromConn(conn); err != nil {
			t.Error(err)
		}
		got = strings.TrimSpace(got)
		fmt.Println("Server:", got)
		if strings.HasSuffix(got, "Bye.") {
			break
		}
		if len(lines) == 0 {
			break
		}
		fmt.Fprintln(conn, lines[0])
		lines = lines[1:]
	}
	conn.Close()
}

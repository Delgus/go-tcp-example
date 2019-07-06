package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	errCreateListener     = "example.tcp-server: can not create listener"
	errCloseListener      = "example.tcp-server: can not close listener"
	errAcceptConnection   = "example.tcp-server: can not accept connection"
	errCloseConnection    = "example.tcp-server: can not close connection"
	errReadFromConnection = "example.tcp-server: can not read from connection"
	errSendToConnection   = "example.tcp-server: can not read from connection"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "127.0.0.1:3333", "server address")
	flag.Parse()

	fmt.Printf("start delgusql-tcp server on %s\n", address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("%s: %v\n", errCreateListener, err)
		os.Exit(1)
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("%s: %v\n", errCloseListener, err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("%s: %v\n", errAcceptConnection, err)
			continue
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("%s %v\n", errCloseConnection, err)
		}
	}()
	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Printf("%s: %v\n", errReadFromConnection, err)
		}

		// send to socket
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(command + "\n")
		if err != nil {
			fmt.Printf("%s: %v\n", errSendToConnection, err)
		}
		err = writer.Flush()
		if err != nil {
			break
		}
	}
}

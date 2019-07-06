package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	errNotFoundConnection = "example.tcp-client: can not connect to server"
	errReadFromInput      = "example.tcp-client: can not read from stdin"
	errSendToConnection   = "example.tcp-client: can not send command to server"
	errGetResponse        = "example.tcp-client: can not get response from server"
)

func main() {

	var address string
	flag.StringVar(&address, "address", "127.0.0.1:3333", "server address")
	flag.Parse()

	// connect to this socket
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%s: %v\n", errNotFoundConnection, err)
		return
	}

	for {
		fmt.Print("tcp.console>>>")

		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s: %v\n", errReadFromInput, err)
			break
		}

		// send to socket
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(text + "\n")
		if err != nil {
			fmt.Printf("%s: %v\n", errSendToConnection, err)
		}
		err = writer.Flush()
		if err != nil {
			break
		}

		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("%s: %v\n", errGetResponse, err)
		}
		fmt.Print(message)
	}
}

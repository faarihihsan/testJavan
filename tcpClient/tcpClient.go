package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8088"
	TYPE = "tcp"
)

func main() {
	for true {
		tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP(TYPE, nil, tcpServer)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}

		var command string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		command = strings.Replace(text, "\n", "", -1)

		_, err = conn.Write([]byte(command))
		if err != nil {
			println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 4096)
		//var received []byte
		_, err = conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			os.Exit(1)
		}

		println("=>", string(received))

		conn.Close()
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func addIpAddress(ip string, port string) {
	_, err := dbClient.Exec("INSERT INTO ip_address(ip, port) VALUES ($1, $2)",
		ip, port)
	if err != nil {
		fmt.Printf("Error while writing ip to db: %v. IP: %v Port: %v\n", err.Error(), ip, port)
	}
}

func sendNotification() {
	address := getIps()

	for _, addr := range address {
		ip := "localhost"
		if addr.Ip != "::1" {
			ip = addr.Ip
		}
		tcpServer, err := net.ResolveTCPAddr("tcp", ip+":"+addr.Port)

		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			continue
		}

		conn, err := net.DialTCP("tcp", nil, tcpServer)
		if err != nil {
			println("Dial failed:", err.Error())
			continue
		}

		_, err = conn.Write([]byte("Data Aset Keluarga Has Changed"))
		if err != nil {
			println("Write data failed:", err.Error())
		}

		// buffer to get data
		received := make([]byte, 4096)
		//var received []byte
		_, err = conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
		}

		println("=>", string(received))

		conn.Close()
	}
}

func getIps() []Adress {
	row := dbClient.QueryRow("select array_to_json(array_agg(to_json(a))) from (select ip, port from ip_address) a;")

	var str []byte
	err := row.Scan(&str)
	if err != nil {
		fmt.Printf("error scanning rows: %v\n", err.Error())
	}

	var res []Adress
	err = json.Unmarshal(str, &res)
	if err != nil {
		fmt.Printf("error scanning rows: %v", err.Error())
	}

	return res
}

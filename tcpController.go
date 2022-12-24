package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

func tcpController(wg *sync.WaitGroup) {
	fmt.Println(fmt.Sprintf("Listening TCP protocol on port %v", tcp_port))
	listen, err := net.Listen("tcp", fmt.Sprintf("%v:%v", tcp_host, tcp_port))
	if err != nil {
		log.Fatal(err)
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listen)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}

	wg.Done()
}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	response := "Success!"

	buffer = bytes.Trim(buffer, "\x00")
	command := strings.Split(string(buffer), " ")
	if command[0] == "GET" {
		response, err = getListKeluargaTcp()
	} else if command[0] == "ADD" {
		if command[1] == "KELUARGA" {
			err = AddKeluargaTcp(command)
		} else if command[1] == "ASET" {
			err = AddAsetTcp(command)
		} else if command[1] == "ASET_KELUARGA" {
			err = AddAsetKeluargaTcp(command)
		} else {
			err = errors.New("command idx 1 is false")
		}
	} else if command[0] == "UPDATE" {
		if command[1] == "KELUARGA" {
			err = UpdateKeluargaTcp(command)
		} else if command[1] == "ASET" {
			err = UpdateAsetTcp(command)
		} else {
			err = errors.New("command idx 1 is false")
		}
	} else if command[0] == "DELETE" {
		if command[1] == "KELUARGA" {
			err = DeleteKeluargaTcp(command)
		} else if command[1] == "ASET" {
			err = DeleteAsetTcp(command)
		} else if command[1] == "ASET_KELUARGA" {
			err = DeleteAsetKeluargaTcp(command)
		} else {
			err = errors.New("command idx 1 is false")
		}
	} else {
		err = errors.New("command idx 0 is false")
	}

	if err != nil {
		response = fmt.Sprintf("Command failed : %v", err)
	}

	conn.Write([]byte(response))

	conn.Close()
}

func getListKeluargaTcp() (string, error) {
	response, err := GetListKeluargaImpl()
	if err != nil {
		return "", err
	}

	out, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func AddKeluargaTcp(command []string) error {
	size := len(command)

	nama := ""
	for i := 2; i < size-1; i++ {
		nama = nama + command[i] + " "
	}

	parentVal := 0
	if command[size-1] != "null" {
		val, err := strconv.Atoi(command[size-1])
		if err != nil {
			return err
		}
		parentVal = int(val)
	}

	err := AddKeluargaImpl(KeluargaPayload{
		Nama:   nama,
		Parent: parentVal,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateKeluargaTcp(command []string) error {
	size := len(command)

	nama := ""
	for i := 3; i < size-1; i++ {
		nama = nama + command[i] + " "
	}

	parentVal := 0
	if command[size-1] != "null" {
		val, err := strconv.Atoi(command[size-1])
		if err != nil {
			return err
		}
		parentVal = int(val)
	}

	id, err := strconv.Atoi(command[2])
	if err != nil {
		return err
	}

	err = UpdateKeluargaImpl(KeluargaPayload{
		Id:     id,
		Nama:   nama,
		Parent: parentVal,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteKeluargaTcp(command []string) error {
	id, err := strconv.Atoi(command[2])
	if err != nil {
		return err
	}

	err = DeleteKeluargaImpl(KeluargaPayload{
		Id: id,
	})
	if err != nil {
		return err
	}

	return nil
}

func AddAsetTcp(command []string) error {
	nama := ""
	for i := 2; i < len(command)-1; i++ {
		nama = nama + command[i] + " "
	}
	nama = nama + command[len(command)-1]

	err := AddAsetImpl(AsetPayload{
		Nama: nama,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateAsetTcp(command []string) error {
	nama := ""
	for i := 3; i < len(command)-1; i++ {
		nama = nama + command[i] + " "
	}
	nama = nama + command[len(command)-1]

	id, err := strconv.Atoi(command[2])
	if err != nil {
		return err
	}

	err = UpdateAsetImpl(AsetPayload{
		Id:   id,
		Nama: nama,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteAsetTcp(command []string) error {
	id, err := strconv.Atoi(command[2])
	if err != nil {
		return err
	}

	err = DeleteAsetImpl(AsetPayload{
		Id: id,
	})
	if err != nil {
		return err
	}

	return nil
}

func AddAsetKeluargaTcp(command []string) error {
	idKeluarga := 0
	idKeluarga, err := strconv.Atoi(command[2])
	idAset := 0
	idAset, err = strconv.Atoi(command[3])
	if err != nil {
		return err
	}

	err = AddAsetKeluargaImpl(AsetKeluargaPayload{
		IdKeluarga: idKeluarga,
		IdAset:     idAset,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteAsetKeluargaTcp(command []string) error {
	id, err := strconv.Atoi(command[2])
	if err != nil {
		return err
	}

	err = DeleteAsetKeluargaImpl(AsetKeluargaPayload{
		Id: id,
	})
	if err != nil {
		return err
	}

	return nil
}

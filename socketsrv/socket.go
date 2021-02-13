/*
socket.go provides a simple TCP socket server
for incoming messages
Copyright © 2021 Andreas Kreisig
Released under the terms of the MIT License
*/

package socketsrv

import (
	//"fmt"
	"log"
	"net"
	//"os"
)

func RfmonListener() {

	service := ":56565"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	errHandler(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	errHandler(err)

	for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	
	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		errHandler(err)
	
		data := buf[:size]
		log.Println("Read new data from connection", data)
		conn.Write(data) // for now it's a simple echo server
		}
	}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
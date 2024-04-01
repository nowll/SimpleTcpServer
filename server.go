package main

import (
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:988")

	handler.ErrorHandler(err)

	defer listen.Close()

	for {
		conn, err := listen.Accept()

		handler.ErrorHandler(err)

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	timeoutDuration := 5 * time.Second

	for {

		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		payload, err := types.Decode(conn)

		handler.ErrorHandler(err)

		fmt.Println("Message :  ", string(payload.Bytes()))

		var w types.Binary

		w = types.Binary("Message Received : " + string(payload.Bytes()))

		_, err = w.WriteTo(conn)

		handler.ErrorHandler(err)
	}
}

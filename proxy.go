package main

import (
	"io"
	"main/handler"
	"net"
)

func main() {
	listen, err := net.Listen("tcp","localhost:767")

	handler.ErrorHandler(err)

	for{
		conn,err := listen.Accept()

		handler.ErrorHandler(err)

		go handlerServer(conn)
	}
}

func clientHandler(sender io.Writer,reader io.Reader) error{

	fromReader, isFromReader := reader.(io.Writer)
	toSender,isToSender := sender.(io.Reader)


	if isFromReader && isToSender{
		go func(){
			_, err := io.Copy(fromReader,toSender)

			handler.ErrorHandler(err)

			return
		}()	
	}

	_,err := io.Copy(sender,reader)

	handler.ErrorHandler(err)

	return err
}


func handlerServer(from net.Conn){
	dial,err := net.Dial("tcp","localhost:988")

	handler.ErrorHandler(err)

	err = clientHandler(from,dial)

	handler.ErrorHandler(err)
}
package main

import (
	"bufio"
	"fmt"
	"main/handler"
	"main/types"
	"net"
	"os"
	"time"
)

var choose int

func printMenu() {
	fmt.Println("1. Send Message.")
	fmt.Println("2. Exit.")
	fmt.Print(">>  ")
	fmt.Scanf("%d\n", &choose)
}

func main() {
	dial, err := net.Dial("tcp", "localhost:767")

	handler.ErrorHandler(err)

	defer dial.Close()

	sc := bufio.NewScanner(os.Stdin)

	var text string

	for {
		printMenu()
		switch choose {
		case 1:

			for {
				fmt.Println("Send Message : ")
				sc.Scan()
				text = sc.Text()
				if len(text) > 4 {
					break
				} else {
					fmt.Println("Char not long enough")
				}

			}
			data := types.Binary(text)
			_, err := data.WriteTo(dial)

			handler.ErrorHandler(err)

			dial.SetReadDeadline(time.Now().Add(5 * time.Second))

			p, err := types.Decode(dial)

			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					fmt.Println("timeout")
					return
				} else {
					handler.ErrorHandler(err)
				}
			}

			fmt.Println(string(p.Bytes()))
		case 2:
			return
		}

	}
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Connect to server
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Ask for a file
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scan(&fileName)

	//Sends file text to server
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		buffer := make([]byte, 1024)
		_, err = file.Read(buffer)
		if err != nil {
			break
		}
		conn.Write(buffer)
	}

	//Receives answer from server
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	

	//Close connection
	conn.Close()

}

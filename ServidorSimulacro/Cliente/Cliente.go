package main

import (
	"fmt"
	"net"
)

func main() {
	// Se conecta al servidor
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Pide input al usuario
	fmt.Print("Escriba 'Hora' para saber la hora actual o 'Fecha' para saber la fecha actual: ")
	var input string
	fmt.Scanln(&input)

	// Se envia el mensaje al servidor
	conn.Write([]byte(input))
	
	// Se recibe la respuesta del servidor
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))
	
	// Se cierra la conexion
	conn.Close()
}
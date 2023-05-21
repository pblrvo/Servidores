package main

import (
	"fmt"
	"net"
	"time"
	"strings"
)

func main() {
	//Se crea direccion tcp para el servidor
	listenAddress, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Se crea el socket para el servidor
	listener, err := net.ListenTCP("tcp", listenAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Se aceptan conexiones
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Se crea un hilo para cada conexion
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	//Se crea un buffer para recibir los datos
	buffer := make([]byte, 1024)

	//Se lee el mensaje
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Se convierte el mensaje a string
	palabra := string(buffer[:n])

	//Se chequea la palabra
	resultado := chequearPalabra(palabra)

	//Se envia el resultado
	conn.Write([]byte(resultado))

}

func chequearPalabra(palabra string) string {
	//Convert to lower case
	palabra = strings.ToLower(palabra)
	if palabra == "hora"{
		return devolverHora()
	} else if palabra == "fecha"{
		return devolverFecha()
	} else {
		return devolverError()
	}
}

func devolverHora() string {
	return string(time.Now().Format("15:04:05"))
}

func devolverFecha() string {
	return  string(time.Now().Format("02/01/2006"))
}

func devolverError() string {
	return "Error"
}



package main

import (
	"log"
	"net"
)

func main() {
	go server()
	go client()

	for {

	}
}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer) // Read seq

	println("Server")
	println((buffer[0]))

	buffer[0] += 1

	conn.Write(buffer)      // Send seq + 1
	conn.Write([]byte{200}) // Send ack

	conn.Read(buffer) // Read seq + 1

	println("Server")
	println((buffer[0]))

	conn.Read(buffer) // Read ack + 1

	println((buffer[0]))
}

func client() {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	dial.Write([]byte{100}) // Send seq

	buffer := make([]byte, 1024)

	dial.Read(buffer) // Read seq + 1

	println("Client")
	println((buffer[0]))

	dial.Write(buffer) // Write seq + 1
	dial.Read(buffer)  // Read ack

	println((buffer[0]))

	buffer[0] += 1

	dial.Write(buffer) // Send ack + 1
}

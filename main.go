package main

import (
	"fmt"
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
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
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
	seq := make([]byte, 1)
	ack := make([]byte, 1)
	ack[0] = 200

	conn.Read(seq) // Read seq

	//Adds 1 to the seq recived
	seq[0] += 1

	//Send the seq + 1 and ack
	conn.Write(seq) // Sends seq + 1
	conn.Write(ack) // Send ack

	//Reads the ack + 1 and seq + 1
	seqRecived := make([]byte, 1)
	ackRecived := make([]byte, 1)
	conn.Read(seqRecived)
	conn.Read(ackRecived)

	//Test that they are correct
	if seqRecived[0] == seq[0] && ackRecived[0] == (ack[0]+1) {
		fmt.Println("Correct seq and ack recived")
	} else {
		fmt.Println("Not correct")
	}
}

func client() {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//Makes the byte to send
	seqSend := make([]byte, 1)
	seqSend[0] = 100

	// Send seq
	dial.Write(seqSend)

	//Make recivers
	ackReceived := make([]byte, 1)
	seqRecived := make([]byte, 1)

	//Reads recives values and prints them
	dial.Read(seqRecived)  // Reads seq + 1
	dial.Read(ackReceived) // Reads ack

	//Checks that the correct value is recived
	if seqRecived[0] == (seqSend[0] + 1) {
		fmt.Println("Correct seq recived")
	} else {
		fmt.Println("Not correct")
	}
	//Adds 1 to the ack
	ackReceived[0] += 1
	//Sends back the ack + 1 and recived to establish the connection
	dial.Write(seqRecived)
	dial.Write(ackReceived)
}
